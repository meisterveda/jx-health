package health

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jenkins-x-plugins/jx-health/pkg/health/lookup"

	"sigs.k8s.io/yaml"

	"github.com/Comcast/kuberhealthy/v2/pkg/khstatecrd"
	"github.com/stretchr/testify/assert"
)

func TestHealthOptions_GetJenkinsXTable(t *testing.T) {
	o := Options{}

	tests := []struct {
		name string
		info bool
		want [][]string
	}{
		{name: "kh_defaults_ok", want: [][]string{
			{"daemonset", "kh-test", "OK", ""},
			{"deployment", "kh-test", "OK", ""},
			{"dns-status-internal", "kh-test", "OK", ""},
		}},
		{name: "kh_defaults_one_fail", want: [][]string{
			{"daemonset", "kh-test", "OK", ""},
			{"deployment", "kh-test", "ERROR", "something bad"},
			{"", "", "", "happened again"},
			{"dns-status-internal", "kh-test", "OK", ""},
		}},
		{name: "kh_info", info: true, want: [][]string{
			{"daemonset", "kh-test", "OK", "", "https://github.com/Comcast/kuberhealthy/blob/230c4f1/cmd/daemonset-check/README.md"},
			{"deployment", "kh-test", "ERROR", "something bad", "https://github.com/Comcast/kuberhealthy/blob/230c4f1/cmd/deployment-check/README.md"},
			{"", "", "", "happened again"},
			{"dns-status-internal", "kh-test", "OK", "", "https://github.com/Comcast/kuberhealthy/blob/230c4f1/cmd/dns-resolution-check/README.md"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.info {
				o.Info = true
				var err error
				o.InfoData, err = lookup.NewLookupData()
				assert.NoError(t, err, "failed to lookup test data")
			}

			states := loadTestStates(t, tt.name)

			got := o.populateTable(states)

			if !reflect.DeepEqual(got, tt.want) {
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
