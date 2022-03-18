package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"civetAdminCenter"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(PermissionMImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("PermissionMImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(civetAdminCenter.PermissionM)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".PermissionMObj")
	
	// Run application
	tars.Run()
}
