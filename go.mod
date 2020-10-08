module github.com/jenkins-x-plugins/jx-health

require (
	github.com/Comcast/kuberhealthy/v2 v2.2.1-0.20201008180926-54448ab4f4c8
	github.com/cpuguy83/go-md2man v1.0.10
	github.com/go-bindata/go-bindata v3.1.2+incompatible // indirect
	github.com/jenkins-x-plugins/jx-scm v0.0.4
	github.com/jenkins-x/go-scm v1.5.175
	github.com/jenkins-x/jx-helpers/v3 v3.0.4
	github.com/jenkins-x/jx-kube-client/v3 v3.0.0
	github.com/jenkins-x/jx-logging/v3 v3.0.2
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	gopkg.in/src-d/go-git.v4 v4.13.1
	gopkg.in/yaml.v1 v1.0.0-20140924161607-9f9df34309c0
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	sigs.k8s.io/yaml v1.2.0
)

go 1.15
