package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/admin/system/tars-protocol/adminApp"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(SystemImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("SystemImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(adminApp.System)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".SystemObj")

	// Run application
	tars.Run()
}
