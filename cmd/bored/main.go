package main

import (
	"flag"
	"os"
)

var (
	configFile  = flag.String("config", "bored.yaml", "relative path to config file")
	versionFlag = flag.Bool("version", false, "version")
)

func main() {
	flag.Parse()

	app, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}

	os.Exit(0)
}
