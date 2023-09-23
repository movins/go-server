package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/admin/common/servants"
	"github.com/admin/common/tars-protocol/adminApp"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(servants.CommonImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("CommonImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(adminApp.Common)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".CommonObj")

	// Run application
	tars.Run()
}
