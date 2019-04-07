package main

import (
	"fmt"
	"os"

	"github.com/Pluslab/gaia/server/config"
	"github.com/Pluslab/gaia/server/server"
)

func main() {
	// call config.Load() before start up
	err := config.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s, err := server.NewEchoServer(config.Port)
	if err != nil {
		fmt.Println("port not specified")
		fmt.Println(err)
		os.Exit(1)
	}

	s.Run()
}
