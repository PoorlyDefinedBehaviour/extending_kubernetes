package main

import (
	"flag"

	"extend-k8s.io/device-plugin/pkg"
	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
)

func main() {
	flag.Parse()
	manager := dpm.NewManager(pkg.Lister{})
	manager.Run()
}
