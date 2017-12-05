package main

import (
	"os"

	logs "github.com/appscode/go/log/golog"
	"github.com/pharmer/openstack-cloud-controller-manager/cmds"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmds.NewRootCmd(Version).Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
