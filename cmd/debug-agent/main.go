package main

import (
	"flag"
	"log"
	"github.com/jamestgrant/kubectl-debug/pkg/debug-agent"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var configFile string
	flag.StringVar(&configFile, "config.file", "", "Config file location.")
	flag.Parse()

	config, err := debug-agent.LoadFile(configFile)
	if err != nil {
		log.Fatalf("error reading config %v", err)
	}

	server, err := debug-agent.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped, see you next time!")
}
