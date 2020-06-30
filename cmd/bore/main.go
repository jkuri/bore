package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jkuri/bore/client"
)

var help = `
Usage: bore [options]

Options:

-s, SSH server remote host (default: bore.jan.local)

-p, SSH server remote port (default: 2200)

-ls, Local HTTP server host (default: localhost)

-lp, Local HTTP server port (default: 7500)

-a, Keep tunnel connection alive (default: true)

Read more:
	https://github.com/jkuri/bore
`

var (
	remoteServer = flag.String("s", "bore.jan.local", "")
	remotePort   = flag.Int("p", 2200, "")
	localServer  = flag.String("ls", "localhost", "")
	localPort    = flag.Int("lp", 80, "")
	keepAlive    = flag.Bool("a", true, "")
)

func main() {
	flag.Usage = func() {
		fmt.Print(help)
		os.Exit(1)
	}
	flag.Parse()

	client := client.NewBoreClient(client.Config{
		RemoteServer: *remoteServer,
		RemotePort:   *remotePort,
		LocalServer:  *localServer,
		LocalPort:    *localPort,
		KeepAlive:    *keepAlive,
	})

	if err := client.Run(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
