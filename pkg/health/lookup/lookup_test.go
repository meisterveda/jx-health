package lookup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLookupData(t *testing.T) {

	tests := []struct {
		name    string
		want    LoopkupData
		wantErr bool
	}{
		{name: "simple", want: LoopkupData{
			Info: map[string]string{
				"daemonset":  "https://github.com/Comcast/kuberhealthy/blob/230c4f1/cmd/daemonset-check/README.md",
				"deployment": "https://github.com/Comcast/kuberhealthy/blob/230c4f1/cmd/deployment-check/README.md",
			}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLookupData()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLookupData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := range tt.want.Info {
				assert.Equal(t, got.Info[i], tt.want.Info[i], "NewLookupData() got = %v, want %v", got.Info[i], tt.want.Info[i])
			}
		})
	}
}
