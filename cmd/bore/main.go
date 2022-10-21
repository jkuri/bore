package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jkuri/bore/client"
	"github.com/jkuri/bore/internal/version"
)

var help = `
Usage: bore [options]

Options:

-s, SSH server remote host (default: bore.services)

-p, SSH server remote port (default: 2200)

-ls, Local HTTP server host (default: localhost)

-lp, Local HTTP server port (default: 7500)

-bp, Remote TCP bind port, (default: 0 (random))

-id, ID to use when generating URL (default: "" (random))

-a, Keep tunnel connection alive (default: true)

-r, Auto-reconnect if connection failed (default: false)

-version, prints bore version and build info

Read more:
	https://github.com/jkuri/bore
`

var (
	remoteServer  = flag.String("s", "bore.services", "")
	remotePort    = flag.Int("p", 2200, "")
	localServer   = flag.String("ls", "localhost", "")
	localPort     = flag.Int("lp", 80, "")
	bindPort      = flag.Int("bp", 0, "")
	id            = flag.String("id", "", "")
	keepAlive     = flag.Bool("a", true, "")
	autoReconnect = flag.Bool("r", false, "")
	versionFlag   = flag.Bool("version", false, "version")
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

connect:
	if err := client.Run(); err != nil {
		if !*autoReconnect {
			log.Fatal(err)
		}
		log.Println("connection failed due: ", err.Error(), "reconnecting in 5s...")
		time.Sleep(time.Second * 5)
		goto connect
	}

	os.Exit(0)
}
