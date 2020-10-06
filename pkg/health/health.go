package health

import (
	"strconv"
	"strings"

	"github.com/Comcast/kuberhealthy/v2/pkg/khstatecrd"

	"github.com/jenkins-x-plugins/jx-health/pkg/options"
	"github.com/jenkins-x/jx-helpers/v3/pkg/table"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const resourceStates = "khstates"

// HealthOptions common CLI arguments for working with health
type HealthOptions struct {
	options.KHCheckOptions
}

func (o HealthOptions) GetJenkinsXTable(ns string) (table.Table, error) {
	result := table.Table{}

	err := o.KHCheckOptions.Validate()
	if err != nil {
		return result, errors.Wrapf(err, "failed to validate KHCheckOptions")
	}

	// get a list of all Kuberhealthy states
	states, err := o.KHCheckOptions.StateClient.List(metav1.ListOptions{}, resourceStates, ns)
	if err != nil {
		return table.Table{}, errors.Wrapf(err, "failed to list jenkins x states in namespace %s", ns)
	}

	o.makeTable(&result, states)

	return result, nil
}

func (o HealthOptions) makeTable(result *table.Table, checks *khstatecrd.KuberhealthyStateList) {
	// add table headers
	result.AddRow("Check Name", "OK", "Error Message")

	// add Kuberhealthy check results to the table
	for _, check := range checks.Items {
		result.AddRow(check.Name, strconv.FormatBool(check.Spec.OK), strings.Join(check.Spec.Errors, "\n"))
	}
}
