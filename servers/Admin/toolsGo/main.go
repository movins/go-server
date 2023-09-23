package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/admin/tools/tars-protocol/adminApp"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(ToolsImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("ToolsImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(adminApp.Tools)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".ToolsObj")

	// Run application
	tars.Run()
}
