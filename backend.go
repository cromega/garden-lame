package main

import (
	"github.com/cloudfoundry-incubator/garden"
	"github.com/pivotal-golang/lager"
	"time"
)

type lameBackend struct {
	garden.Client
	logger lager.Logger
}

func (b *lameBackend) Start() error {
	b.logger.Info("backend-start")
	return nil
}

func (b *lameBackend) Stop() {
	b.logger.Info("backend-stop")
}

func (b *lameBackend) GraceTime(garden.Container) time.Duration {
	return 0
}
