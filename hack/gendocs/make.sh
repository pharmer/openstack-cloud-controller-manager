#!/usr/bin/env bash

pushd $GOPATH/src/github.com/pharmer/openstack-cloud-controller-manager/hack/gendocs
go run main.go
popd
