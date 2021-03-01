module github.com/jenkins-x-plugins/jx-health

require (
	github.com/Comcast/kuberhealthy/v2 v2.2.1-0.20201008180926-54448ab4f4c8
	github.com/alecthomas/assert v0.0.0-20170929043011-405dbfeb8e38
	github.com/alecthomas/colour v0.1.0 // indirect
	github.com/alecthomas/repr v0.0.0-20201103221029-55c485bd663f // indirect
	github.com/cpuguy83/go-md2man v1.0.10
	github.com/jenkins-x/jx-helpers/v3 v3.0.84
	github.com/jenkins-x/jx-kube-client/v3 v3.0.2
	github.com/jenkins-x/jx-logging/v3 v3.0.3
	github.com/liggitt/tabwriter v0.0.0-20181228230101-89fcab3d43de
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kubernetes v1.14.7
	sigs.k8s.io/yaml v1.2.0
)

replace (
	k8s.io/api => k8s.io/api v0.20.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.2
	k8s.io/client-go => k8s.io/client-go v0.20.2
)

// lets use a PR of kuberhealthy until this PR is merged: https://github.com/Comcast/kuberhealthy/pull/758
replace github.com/Comcast/kuberhealthy/v2 => github.com/jstrachan/kuberhealthy/v2 v2.3.2-0.20201211103805-042693101caa

go 1.15
