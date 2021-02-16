package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/pankratsdarya/golevelone/configuration/config"
)

var (
	mainPort   = flag.Int("main-port", 0, "PORT configuration")
	someAppID  = flag.String("AppID", "ID", "App ID configuration. Starts with ID.")
	someAppKEY = flag.String("AppKEY", "KEY", "App Key configuration. Starts with KEY.")
)

func main() {
	flag.Parse()
	fullConfig, _ := config.GetConfig()

	// Input validation
	if *mainPort < 0 {
		log.Fatal("Invalid flag: main-port.")
	}
	if !strings.HasPrefix(*someAppID, "ID") {
		log.Fatal("Invalid flag: AppID.")
	}
	if !strings.HasPrefix(*someAppKEY, "KEY") {
		log.Fatal("Invalid flag: AppKEY.")
	}

	// Flags has higher priority. Set configuration as in flags
	if fullConfig != nil {
		if *mainPort != 0 {
			fullConfig.MainPort = *mainPort
		}
	}
	if len(*someAppID) > 2 {
		fullConfig.SomeAppID = *someAppID
	}
	if len(*someAppKEY) > 3 {
		fullConfig.SomeAppKEY = *someAppKEY
	}

	fmt.Println(fullConfig)

}
