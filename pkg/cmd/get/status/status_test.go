package status

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert"
)

func TestOptions_getNamespace(t *testing.T) {
	// getting the current namespace is found from a local kube config file
	err := os.Setenv("KUBECONFIG", filepath.Join("test_data", "test-config"))
	assert.NoError(t, err)

	type fields struct {
		AllNamespaces bool
		Namespace     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{name: "all_namespaces",
			fields: fields{
				AllNamespaces: true,
			}, want: "", wantErr: false},
		{name: "all_namespaces_priority",
			fields: fields{
				AllNamespaces: true,
				Namespace:     "bar",
			}, want: "", wantErr: false},
		{name: "specific_ns",
			fields: fields{
				AllNamespaces: false,
				Namespace:     "bar",
			}, want: "bar", wantErr: false},
		{name: "current_ns",
			fields: fields{
				AllNamespaces: false,
			}, want: "cheese", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				AllNamespaces: tt.fields.AllNamespaces,
				Namespace:     tt.fields.Namespace,
			}
			got, err := o.getNamespace()
			if (err != nil) != tt.wantErr {
				t.Errorf("getNamespace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getNamespace() got = %v, want %v", got, tt.want)
			}
		})
	}
}
