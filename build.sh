#! /bin/sh
set -e

GOPATH=$(go env GOPATH)
SRC=$GOPATH/src
BIN=$GOPATH/bin
ROOT=$GOPATH
REPO_ROOT=$GOPATH/src/github.com/aerokite/kubemerge

export CGO_ENABLED=0

mkdir -p $REPO_ROOT/bin
go build -v -o $REPO_ROOT/bin/kubemerge $REPO_ROOT/main.go
