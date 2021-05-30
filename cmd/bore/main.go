package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jkuri/bore/client"
	"github.com/jkuri/bore/internal/version"
)

var help = `
Usage: bore [options]

Options:

-s, SSH server remote host (default: bore.network)

-p, SSH server remote port (default: 2200)

-ls, Local HTTP server host (default: localhost)

-lp, Local HTTP server port (default: 7500)

-bp, Remote TCP bind port, (default: 0 (random))

-id, ID to use when generating URL (default: "" (random))

-a, Keep tunnel connection alive (default: true)

-version, prints bore version and build info

Read more:
	https://github.com/jkuri/bore
`

var (
	remoteServer = flag.String("s", "bore.network", "")
	remotePort   = flag.Int("p", 2200, "")
	localServer  = flag.String("ls", "localhost", "")
	localPort    = flag.Int("lp", 80, "")
	bindPort     = flag.Int("bp", 0, "")
	id           = flag.String("id", "", "")
	keepAlive    = flag.Bool("a", true, "")
	versionFlag  = flag.Bool("version", false, "version")
)

func main() {
	flag.Usage = func() {
		fmt.Print(help)
		os.Exit(1)
	}
	flag.Parse()

	if *versionFlag {
		fmt.Printf("%s\n", version.GenerateBuildVersionString())
		os.Exit(0)
	}

	client := client.NewBoreClient(client.Config{
		RemoteServer: *remoteServer,
		RemotePort:   *remotePort,
		LocalServer:  *localServer,
		LocalPort:    *localPort,
		BindPort:     *bindPort,
		ID:           *id,
		KeepAlive:    *keepAlive,
	})

	if err := client.Run(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
