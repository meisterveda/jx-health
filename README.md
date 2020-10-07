# jx-health

[![Documentation](https://godoc.org/github.com/jenkins-x-plugins/jx-health?status.svg)](https://pkg.go.dev/mod/github.com/jenkins-x-plugins/jx-health)
[![Go Report Card](https://goreportcard.com/badge/github.com/jenkins-x-plugins/jx-health)](https://goreportcard.com/report/github.com/jenkins-x-plugins/jx-health)
[![Releases](https://img.shields.io/github/release-pre/jenkins-x/jx-health.svg)](https://github.com/jenkins-x-plugins/jx-health/releases)
[![LICENSE](https://img.shields.io/github/license/jenkins-x/jx-health.svg)](https://github.com/jenkins-x-plugins/jx-health/blob/master/LICENSE)
[![Slack Status](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://slack.k8s.io/)

jx-health is a small command line tool working with health statuses from [Kuberhealthy](https://github.com/Comcast/kuberhealthy)

Using Kuberhealthy and custom checks we are able to report of the health of a Jenkins X installation by only reading the Kuberhealthy state custom resources.  This is good for user RBAC restricted environments as the Kuebrhealthy checks run with a Kubernetes service account to validate things like secrets, without revealing any sensitive data and report errors when a user may not have access.   
## Getting Started

Download the [jx-health binary](https://github.com/jenkins-x-plugins/jx-health/releases) for your operating system and add it to your `$PATH`.

## Commands

See the [jx-health command reference](docs/cmd/jx-health.md#see-also)

## Developing

Golang 1.15

You can now build this repository using your local modifications and try the locally built binary in `build/jx-health` or run the unit tests via `make test`