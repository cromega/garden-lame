package main

import (
	"github.com/cloudfoundry-incubator/garden/server"
	"github.com/pivotal-golang/lager"
	"time"
)

func main() {
	logger := lager.NewLogger("garden-lame")
	backend := lameBackend{Client: &lameClient{}}
	srv := server.New("tcp", ":3000", 5*time.Minute, &backend, logger)
	srv.Start()
}
