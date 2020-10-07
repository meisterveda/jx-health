package status

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func TestOptions_getNamespaces(t *testing.T) {

	// creates two fake namespaces
	client := setupFakeNamespaces()

	// getting the current namespace is found from a local kube config file
	os.Setenv("KUBECONFIG", filepath.Join("test_data", "test-config"))

	type fields struct {
		AllNamespaces bool
		Namespace     string
		KubeClient    kubernetes.Interface
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{name: "all-namespaces", fields: fields{
			AllNamespaces: true,
			Namespace:     "",
			KubeClient:    client,
		}, want: []string{"foo1", "foo2"}, wantErr: false},
		{name: "all-namespaces_with_ns", fields: fields{
			AllNamespaces: true,
			Namespace:     "bad_ns",
			KubeClient:    client,
		}, want: []string{"foo1", "foo2"}, wantErr: false},
		{name: "specific_ns", fields: fields{
			AllNamespaces: false,
			Namespace:     "bar",
			KubeClient:    client,
		}, want: []string{"bar"}, wantErr: false},
		{name: "current_ns", fields: fields{
			AllNamespaces: false,
			Namespace:     "",
			KubeClient:    client,
		}, want: []string{"cheese"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				AllNamespaces: tt.fields.AllNamespaces,
				Namespace:     tt.fields.Namespace,
				KubeClient:    client,
			}
			got, err := o.getNamespaces()
			if (err != nil) != tt.wantErr {
				t.Errorf("getNamespaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNamespaces() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func setupFakeNamespaces() *fake.Clientset {
	ns1 := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "foo1",
		},
	}
	ns2 := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "foo2",
		},
	}
	return fake.NewSimpleClientset(ns1, ns2)
}
