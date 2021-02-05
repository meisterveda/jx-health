module github.com/jenkins-x-plugins/jx-health

require (
	cloud.google.com/go v0.76.0 // indirect
	github.com/Comcast/kuberhealthy/v2 v2.2.1-0.20201008180926-54448ab4f4c8
	github.com/alecthomas/assert v0.0.0-20170929043011-405dbfeb8e38
	github.com/alecthomas/colour v0.1.0 // indirect
	github.com/alecthomas/repr v0.0.0-20201103221029-55c485bd663f // indirect
	github.com/cpuguy83/go-md2man v1.0.10
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.5.4 // indirect
	github.com/gophercloud/gophercloud v0.15.0 // indirect
	github.com/jenkins-x/jx-helpers/v3 v3.0.74
	github.com/jenkins-x/jx-kube-client/v3 v3.0.2
	github.com/jenkins-x/jx-logging/v3 v3.0.3
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/oauth2 v0.0.0-20210201163806-010130855d6c // indirect
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0 // indirect
	k8s.io/klog/v2 v2.5.0 // indirect
	k8s.io/kubernetes v1.14.7
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009 // indirect
	sigs.k8s.io/yaml v1.2.0
)

// lets use a PR of kuberhealthy until this PR is merged: https://github.com/Comcast/kuberhealthy/pull/758
replace github.com/Comcast/kuberhealthy/v2 => github.com/jstrachan/kuberhealthy/v2 v2.3.2-0.20201211103805-042693101caa

go 1.15
