# jx-health

[![Documentation](https://godoc.org/github.com/jenkins-x-plugins/jx-health?status.svg)](https://pkg.go.dev/mod/github.com/jenkins-x-plugins/jx-health)
[![Go Report Card](https://goreportcard.com/badge/github.com/jenkins-x-plugins/jx-health)](https://goreportcard.com/report/github.com/jenkins-x-plugins/jx-health)
[![Releases](https://img.shields.io/github/release-pre/jenkins-x/jx-health.svg)](https://github.com/jenkins-x-plugins/jx-health/releases)
[![LICENSE](https://img.shields.io/github/license/jenkins-x/jx-health.svg)](https://github.com/jenkins-x-plugins/jx-health/blob/master/LICENSE)
[![Slack Status](https://img.shields.io/badge/slack-join_chat-white.svg?logo=slack&style=social)](https://slack.k8s.io/)

jx-health is a small command line tool working with git providers using [go-scm](https://github.com/jenkins-x/go-scm)

## Getting Started

Download the [jx-health binary](https://github.com/jenkins-x-plugins/jx-health/releases) for your operating system and add it to your `$PATH`.

There will be an `app` you can install soon too...

## Commands

See the [jx-health command reference](docs/cmd/jx-health.md#see-also)


## Developing

If you wish to work on a local clone of [go-scm](https://github.com/jenkins-x/go-scm) then:

```bash                  
git clone https://github.com/jenkins-x/go-scm
```                                          

Then in the local `go.mod` file add the following at the end:


``` 
replace github.com/jenkins-x/go-scm  => PathToTheAboveGitClone
```                                                           

You can now build this repository using your local modifications and try the locally built binary in `build/jx-health` or run the unit tests via `make test`