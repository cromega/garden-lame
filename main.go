package main

import (
	"github.com/cloudfoundry-incubator/garden/server"
	"github.com/pivotal-golang/lager"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := lager.NewLogger("garden-lame")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	client := NewLameClient(logger)
	backend := lameBackend{logger: logger, Client: client}
	srv := server.New("tcp", ":3000", 5*time.Minute, &backend, logger)
	srv.Start()

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-exit

	srv.Stop()
}
