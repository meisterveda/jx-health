package options

import (
	"github.com/Comcast/kuberhealthy/v2/pkg/khcheckcrd"
	"github.com/Comcast/kuberhealthy/v2/pkg/khstatecrd"
	"github.com/jenkins-x/jx-kube-client/v3/pkg/kubeclient"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// constants for using the kuberhealthy check CRD
const checkCRDGroup = "comcast.github.io"
const checkCRDVersion = "v1"

// KHCheckOptions common CLI arguments for working with health
type KHCheckOptions struct {
	StateClient *khstatecrd.KuberhealthyStateClient
}

// Validate validates the options and returns the KuberhealthyCheckClient
func (o *KHCheckOptions) Validate() error {

	f := kubeclient.NewFactory()
	cfg, err := f.CreateKubeConfig()
	if err != nil {
		log.Logger().Fatalf("failed to get kubernetes config: %v", err)
	}

	if o.StateClient == nil {
		o.StateClient, err = ClientStateClient(cfg, checkCRDGroup, checkCRDVersion)
		if err != nil {
			return errors.Wrap(err, "failed to create kuberhealthy check client")
		}
	}

	return nil
}

// ClientStateClient creates a rest client to use for interacting with Kuberhealthy state CRDs
func ClientStateClient(cfg *rest.Config, GroupName string, GroupVersion string) (*khstatecrd.KuberhealthyStateClient, error) {
	var err error

	err = khcheckcrd.ConfigureScheme(GroupName, GroupVersion)
	if err != nil {
		return &khstatecrd.KuberhealthyStateClient{}, err
	}

	config := *cfg
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: GroupName, Version: GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.WithoutConversionCodecFactory{CodecFactory: scheme.Codecs}
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	return khstatecrd.CreateClient(client), err
}
