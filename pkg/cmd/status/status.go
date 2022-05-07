package status

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jenkins-x/jx-helpers/v3/pkg/termcolor"
	"github.com/liggitt/tabwriter"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "k8s.io/api/core/v1"

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

const (
	tabwriterMinWidth = 6
	tabwriterWidth    = 4
	tabwriterPadding  = 3
	tabwriterPadChar  = ' '
	tabwriterFlags    = 7
)

var (
	info = termcolor.ColorInfo

	cmdLong = templates.LongDesc(`
		Prints health statuses in a table
`)

	cmdExample = templates.Examples(`
		# prints all health statuses for the current namespace in a table
		%s get status

		# prints all health statuses for all accessible namespace
		%s get status --all-namespaces

		# prints all health statuses for all accessible namespace
		%s get status -A

		# prints all health statuses for a specific namespace
		%s get status --namespace mything

		# watch health statuses
		%s get status --watch
	`)
)

// Options the options for the command
type Options struct {
	HealthOptions         healthopts.Options
	Args                  []string
	Namespace             string
	KuberhealthyNamespace string
	KuberhealthyName      string
	AllNamespaces         bool
	Watch                 bool
	FailIfNoKuberhealthy  bool
	KuberhealthyRunning   bool
	KubeClient            kubernetes.Interface
	cfg                   *rest.Config
}

// NewCmdStatus creates a command object for the command
func NewCmdStatus() (*cobra.Command, *Options) {
	o := &Options{}

	cmd := &cobra.Command{
		Use:     "status",
		Aliases: []string{"statuses"},
		Short:   "Gets health statuses",
		Long:    cmdLong,
		Example: fmt.Sprintf(cmdExample, rootcmd.BinaryName, rootcmd.BinaryName, rootcmd.BinaryName, rootcmd.BinaryName, rootcmd.BinaryName),
		Run: func(cmd *cobra.Command, args []string) {
			o.Args = args
			err := o.Run()
			helper.CheckErr(err)
		},
	}

	cmd.Flags().StringVarP(&o.Namespace, "namespace", "n", "", "namespace to get status checks, defaults to current namespace")
	cmd.Flags().StringVarP(&o.KuberhealthyNamespace, "kuberhealthy-namespace", "", kuberhealthyNamespace, "namespace that kuberhealthy is running")
	cmd.Flags().StringVarP(&o.KuberhealthyName, "kuberhealthy", "", "kuberhealthy", "deployment name of kuberhealthy")

	cmd.Flags().BoolVarP(&o.AllNamespaces, "all-namespaces", "A", false, "if present, list the requested object(s) across all namespaces.\nNamespace in current context is ignored even if specified with --namespace.")
	cmd.Flags().BoolVarP(&o.HealthOptions.Info, "info", "", false, "provide information links for checks")
	cmd.Flags().BoolVarP(&o.FailIfNoKuberhealthy, "fail-if-missing", "f", false, "fail the status check if kuberhealthy is not running")
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
	err = o.verifyKuberhealthyRunning()
	if err != nil {
		return errors.Wrapf(err, "failed to verify Kuberheathy is running")
	}
	return nil
}

func (o *Options) verifyKuberhealthyRunning() error {
	d, err := o.KubeClient.AppsV1().Deployments(o.KuberhealthyNamespace).Get(context.TODO(), o.KuberhealthyName, metav1.GetOptions{})
	if err != nil {
		if o.FailIfNoKuberhealthy {
			return errors.Wrapf(err, "error finding kuberhealthy deployment %s running in the %s namespace, is it installed?", o.KuberhealthyName, o.KuberhealthyNamespace)
		}

		log.Logger().Infof("kuberhealthy is not running in namespace %s with deployment %s", info(o.KuberhealthyName), info(o.KuberhealthyNamespace))
		log.Logger().Infof("for help on installing kuberhealthy see: %s", info("https://jenkins-x.io/v3/admin/setup/health#install"))
		return nil
	}

	if *d.Spec.Replicas != d.Status.ReadyReplicas {
		return errors.Wrapf(err, "not all kuberhealthy pods are running in the %s namespace, expected %d got %d?", kuberhealthyNamespace, d.Spec.Replicas, d.Status.ReadyReplicas)
	}
	o.KuberhealthyRunning = true
	return nil
}

// Run transforms the YAML files
func (o *Options) Run() error {
	err := o.Validate()
	if err != nil {
		return errors.Wrapf(err, "failed to validate status options")
	}

	if !o.KuberhealthyRunning {
		return nil
	}

	// add table headers

	table := tabwriter.NewWriter(os.Stdout, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
	table.Init(os.Stdout, 30, 0, 3, ' ', 7)

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
		err := o.HealthOptions.WriteStatusTable(table, namespace)
		if err != nil {
			return errors.Wrapf(err, "failed to build health table, is Kuberhealthy installed?")
		}
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
