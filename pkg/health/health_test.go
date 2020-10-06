package health

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"sigs.k8s.io/yaml"

	"github.com/Comcast/kuberhealthy/v2/pkg/khstatecrd"
	"github.com/stretchr/testify/assert"

	"github.com/jenkins-x/jx-helpers/v3/pkg/table"
)

func TestHealthOptions_GetJenkinsXTable(t *testing.T) {
	o := HealthOptions{}

	tests := []struct {
		name string
		want [][]string
	}{
		{name: "kh_defaults_ok", want: [][]string{
			{"daemonset", "true", ""},
			{"deployment", "true", ""},
			{"dns-status-internal", "true", ""},
		}},
		{name: "kh_defaults_one_fail", want: [][]string{
			{"daemonset", "true", ""},
			{"deployment", "false", "something bad\nhappened again"},
			{"dns-status-internal", "true", ""},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			expected := table.Table{Rows: [][]string{
				{"Check Name", "OK", "Error Message"},
			}}

			for _, row := range tt.want {
				expected.AddRow(row...)
			}

			states := loadTestStates(t, tt.name)

			got := table.Table{}
			o.makeTable(&got, states)

			if !reflect.DeepEqual(got.Rows, expected.Rows) {
				t.Errorf("GetJenkinsXTable() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// load test states used to build the table
func loadTestStates(t *testing.T, name string) *khstatecrd.KuberhealthyStateList {
	file := filepath.Join("test_data", name, "states.yaml")

	setupData, err := ioutil.ReadFile(file)
	assert.NoError(t, err, "failed to read file")

	states := &khstatecrd.KuberhealthyStateList{}

	err = yaml.Unmarshal(setupData, states)
	assert.NoError(t, err, "failed to unmarshal states yaml")

	return states
}
