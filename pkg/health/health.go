package health

import (
	"sort"

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

		// get matching information link
		informationDetail := o.InfoData.Info[check.Name]

		// depending on if there are errors or how many there are we want to format the table to it is easy to consume
		// Name | Namespace | Status | Error Message        | Info (optional)
		// foo    jx          ok
		// bar    jx          error    first error for bar
		//                             second error for bar
		// cheese jx          ok
		rowEntries := []string{check.Name, check.Namespace, status}
		if len(check.Spec.Errors) == 0 {
			rowEntries = append(rowEntries, "")
			if o.Info {
				rowEntries = append(rowEntries, informationDetail)
			}
			result.AddRow(rowEntries...)
		} else {
			rowEntries = append(rowEntries, check.Spec.Errors[0])
			if o.Info {
				rowEntries = append(rowEntries, informationDetail)
			}
			result.AddRow(rowEntries...)

			// if we have multiple errors lets format the table so all errors appear underneath in the column
			if len(check.Spec.Errors) > 1 {
				for i := 1; i < len(check.Spec.Errors); i++ {
					result.AddRow("", "", "", check.Spec.Errors[i])
				}
			}
		}

	}

}
