package main

import (
	"github.com/cloudfoundry-incubator/garden"
)

type lameProcess struct {
}

func (p *lameProcess) ID() uint32 {
	return 0
}
func (p *lameProcess) Wait() (int, error) {
	return 0, nil
}
func (p *lameProcess) SetTTY(garden.TTYSpec) error {
	return nil
}
func (p *lameProcess) Signal(garden.Signal) error {
	return nil
}
