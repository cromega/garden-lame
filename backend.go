package main

import (
	"fmt"
	"github.com/cloudfoundry-incubator/garden"
	"time"
)

type lameBackend struct {
	garden.Client
}

func (b *lameBackend) Start() error {
	fmt.Println("backend start")
	return nil
}

func (b *lameBackend) Stop() {
	fmt.Println("backend stop")
}

func (b *lameBackend) GraceTime(garden.Container) time.Duration {
	return 0
}
