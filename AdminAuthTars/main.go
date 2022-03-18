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
	imp := new(VerifyImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("VerifyImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(civetAdminCenter.Verify)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".VerifyObj")
	
	// Run application
	tars.Run()
}
