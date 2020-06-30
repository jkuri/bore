package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jkuri/bore/internal/version"
)

var (
	configFile  = flag.String("config", "bore-server.yaml", "relative path to config file")
	versionFlag = flag.Bool("version", false, "version")
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Printf("%s\n", version.GenerateBuildVersionString())
		os.Exit(0)
	}

	app, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}

	os.Exit(0)
}
