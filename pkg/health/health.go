package health

import (
	"sort"
	"strings"

	"github.com/jenkins-x-plugins/jx-health/pkg/health/lookup"

	"github.com/jenkins-x/jx-helpers/v3/pkg/termcolor"

	"github.com/Comcast/kuberhealthy/v2/pkg/khstatecrd"

	"github.com/jenkins-x-plugins/jx-health/pkg/options"
	"github.com/jenkins-x/jx-helpers/v3/pkg/table"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const resourceStates = "khstates"

// Options common CLI arguments for working with health
type Options struct {
	options.KHCheckOptions
	Info     bool
	InfoData lookup.LoopkupData
}

func (o Options) GetJenkinsXTable(result *table.Table, ns string) error {

	err := o.KHCheckOptions.Validate()
	if err != nil {
		return errors.Wrapf(err, "failed to validate KHCheckOptions")
	}

	// get a list of all Kuberhealthy states
	states, err := o.KHCheckOptions.StateClient.List(metav1.ListOptions{}, resourceStates, ns)
	if err != nil {
		return errors.Wrapf(err, "failed to list health states in namespace %s", ns)
	}

	o.populateTable(result, states)

	return nil
}

func (o Options) populateTable(result *table.Table, checks *khstatecrd.KuberhealthyStateList) {

	sort.Slice(checks.Items, func(i, j int) bool {
		return checks.Items[i].Name < checks.Items[j].Name
	})

	// add Kuberhealthy check results to the table
	for _, check := range checks.Items {
		status := termcolor.ColorInfo("OK")
		if !check.Spec.OK {
			status = termcolor.ColorError("ERROR")
		}
		rowEntries := []string{check.Name, check.Namespace, status, termcolor.ColorError(strings.Join(check.Spec.Errors, "\n"))}
		if o.Info {
			informationDetail := o.InfoData.Info[check.Name]

			rowEntries = append(rowEntries, informationDetail)
		}

		result.AddRow(rowEntries...)
	}

}
