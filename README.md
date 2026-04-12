# GLOG

A nice journalctl aggregation logger using CHARM's Log

![GO](https://github.com/lundjrl/glog/actions/workflows/go.yml/badge.svg)

## How to build

1. Make sure you have Go installed
2. Run Go build
3. Run `go list -f '{{.Target}}'`
4. Copy that path (minus /glog) to your GOBIN variable with `go env -w GOBIN=/path/to/your/bin`
5. Run `go install`
6. Run `glog`, enjoy

If stuck, follow the details [here](https://go.dev/doc/tutorial/compile-install)
