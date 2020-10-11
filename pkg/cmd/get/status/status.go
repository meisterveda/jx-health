package status

import (
	"context"
	"fmt"
	"os"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "k8s.io/api/core/v1"

	"github.com/liggitt/tabwriter"

	"k8s.io/kubernetes/pkg/printers"

	"k8s.io/client-go/rest"

	"github.com/jenkins-x/jx-logging/v3/pkg/log"

	"github.com/jenkins-x-plugins/jx-health/pkg/health/lookup"

	"github.com/jenkins-x/jx-kube-client/v3/pkg/kubeclient"

	"k8s.io/client-go/kubernetes"

	healthopts "github.com/jenkins-x-plugins/jx-health/pkg/health"
	"github.com/jenkins-x-plugins/jx-health/pkg/rootcmd"
	"github.com/jenkins-x/jx-helpers/v3/pkg/cobras/helper"
	"github.com/jenkins-x/jx-helpers/v3/pkg/cobras/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const kuberhealthyNamespace = "kuberhealthy"

var (
	cmdLong = templates.LongDesc(`
		Prints health statuses in a table
`)

	cmdExample = templates.Examples(`
		# prints all health statuses for the current namespace in a table
		%s get status

		# prints all health statuses for a specific namespace
		%s get status --namespace

		# prints all health statuses for all accessible namespace
		%s get status --all-namespaces
	`)
)

// LabelOptions the options for the command
type Options struct {
	HealthOptions healthopts.Options
	Args          []string
	AllNamespaces bool
	Watch         bool
	Namespace     string
	KubeClient    kubernetes.Interface
	cfg           *rest.Config
}

// NewCmdStatus creates a command object for the command
func NewCmdStatus() (*cobra.Command, *Options) {
	o := &Options{}

	cmd := &cobra.Command{
		Use:     "status",
		Aliases: []string{"statuses"},
		Short:   "Gets health statuses",
		Long:    cmdLong,
		Example: fmt.Sprintf(cmdExample, rootcmd.BinaryName, rootcmd.BinaryName, rootcmd.BinaryName),
		Run: func(cmd *cobra.Command, args []string) {
			o.Args = args
			err := o.Run()
			helper.CheckErr(err)
		},
	}

	cmd.Flags().BoolVarP(&o.AllNamespaces, "all-namespaces", "A", false, "if present, list the requested object(s) across all namespaces.\nNamespace in current context is ignored even if specified with --namespace.")
	cmd.Flags().BoolVarP(&o.HealthOptions.Info, "info", "", false, "provide information links for checks")
	cmd.Flags().StringVarP(&o.Namespace, "namespace", "n", "", "namespace to get status checks, defaults to current namespace")
	cmd.Flags().BoolVarP(&o.Watch, "watch", "w", false, "after listing/getting the requested object, watch for changes")

	return cmd, o
}

// Validate verifies settings
func (o *Options) Validate() error {
	err := o.HealthOptions.Validate()
	if err != nil {
		return errors.Wrapf(err, "failed to validate health options")
	}

	f := kubeclient.NewFactory()
	o.cfg, err = f.CreateKubeConfig()
	if err != nil {
		return errors.Wrapf(err, "failed to get kubernetes config")
	}

	o.KubeClient, err = kubernetes.NewForConfig(o.cfg)
	if err != nil {
		return errors.Wrapf(err, "error building kubernetes clientset")
	}

	// if user wants information linked to any health check status then lookup the information we have
	if o.HealthOptions.Info {
		o.HealthOptions.InfoData, err = lookup.NewLookupData()
		if err != nil {
			log.Logger().Warnf("unable to lookup health check information: %v", err)
		}
	}

	// check kuberhealthy is installed
	err2, done := o.verifyKuberhealthyRunning(err)
	if done {
		return err2
	}
	return nil
}

func (o *Options) verifyKuberhealthyRunning(err error) (error, bool) {
	d, err := o.KubeClient.AppsV1().Deployments(kuberhealthyNamespace).Get(context.TODO(), "kuberhealthy", metav1.GetOptions{})
	if err != nil {
		return errors.Wrapf(err, "error finding kuberhealthy running in the %s namespace, is it installed?", kuberhealthyNamespace), true
	}

	if *d.Spec.Replicas != d.Status.ReadyReplicas {
		return errors.Wrapf(err, "not all kuberhealthy pods are running in the %s namespace, expected %d got %d?", kuberhealthyNamespace, d.Spec.Replicas, d.Status.ReadyReplicas), true
	}
	return nil, false
}

// Run transforms the YAML files
func (o *Options) Run() error {
	err := o.Validate()
	if err != nil {
		return errors.Wrapf(err, "failed to validate status options")
	}

	// add table headers
	table := printers.GetNewTabWriter(os.Stdout)
	table.Init(os.Stdout, 30, 0, 3, ' ', tabwriter.RememberWidths)

	defaultHeaders := []string{"NAME", "NAMESPACE", "STATUS", "ERROR MESSAGE"}
	if o.HealthOptions.Info {
		defaultHeaders = append(defaultHeaders, "INFORMATION")
	}

	fmt.Fprintln(table, strings.Join(defaultHeaders, "\t"))

	namespace, err := o.getNamespace()
	if err != nil {
		return errors.Wrapf(err, "failed to work out what namespace to use")
	}

	if o.Watch {
		err = o.HealthOptions.WatchStates(table, o.cfg, namespace)
		if err != nil {
			return errors.Wrapf(err, "failed to watch health states")
		}
	} else {
		err := o.HealthOptions.GetJenkinsXTable(table, namespace)
		if err != nil {
			return errors.Wrapf(err, "failed to build health table, is Kuberhealthy installed?")
		}
		table.Flush()
	}

	return nil
}

// decide the namespace to search for kuberhealthy states in this order
// 1. --all-namespaces takes priority
// 2. --namespace when user specifies the namespace
// 3. the default is to search for health statuses in the current namespace
func (o *Options) getNamespace() (string, error) {
	var namespace string
	var err error
	switch {
	case o.AllNamespaces:
		namespace = v1.NamespaceAll
	case o.Namespace != "":
		namespace = o.Namespace
	default:
		namespace, err = kubeclient.CurrentNamespace()
		if err != nil {
			return namespace, errors.Wrapf(err, "failed to find the current namespace")
		}
	}
	return namespace, nil
}
