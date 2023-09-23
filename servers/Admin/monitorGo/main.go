package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/admin/monitor/tars-protocol/adminApp"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(MonitorImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("MonitorImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(adminApp.Monitor)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".MonitorObj")

	// Run application
	tars.Run()
}
