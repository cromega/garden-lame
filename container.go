package main

import (
	"github.com/cloudfoundry-incubator/garden"
	"io"
	"os"
)

type lameContainer struct {
}

func (c *lameContainer) Handle() string {
	return ""
}

func (c *lameContainer) Stop(kill bool) error {
	return nil
}

func (c *lameContainer) Info() (garden.ContainerInfo, error) {
	return garden.ContainerInfo{}, nil
}
func (c *lameContainer) StreamIn(spec garden.StreamInSpec) error {
	return nil
}
func (c *lameContainer) StreamOut(spec garden.StreamOutSpec) (io.ReadCloser, error) {
	return os.Stdout, nil
}

func (c *lameContainer) LimitBandwidth(limits garden.BandwidthLimits) error {
	return nil
}

func (c *lameContainer) CurrentBandwidthLimits() (garden.BandwidthLimits, error) {
	return garden.BandwidthLimits{}, nil
}

func (c *lameContainer) LimitCPU(limits garden.CPULimits) error {
	return nil
}

func (c *lameContainer) CurrentCPULimits() (garden.CPULimits, error) {
	return garden.CPULimits{}, nil
}

func (c *lameContainer) LimitDisk(limits garden.DiskLimits) error {
	return nil
}

func (c *lameContainer) CurrentDiskLimits() (garden.DiskLimits, error) {
	return garden.DiskLimits{}, nil
}

func (c *lameContainer) LimitMemory(limits garden.MemoryLimits) error {
	return nil
}

func (c *lameContainer) CurrentMemoryLimits() (garden.MemoryLimits, error) {
	return garden.MemoryLimits{}, nil
}

func (c *lameContainer) NetIn(hostPort, containerPort uint32) (uint32, uint32, error) {
	return 0, 0, nil
}

func (c *lameContainer) NetOut(netOutRule garden.NetOutRule) error {
	return nil
}

func (c *lameContainer) Run(garden.ProcessSpec, garden.ProcessIO) (garden.Process, error) {
	return &lameProcess{}, nil
}

func (c *lameContainer) Attach(processID uint32, io garden.ProcessIO) (garden.Process, error) {
	return &lameProcess{}, nil
}

func (c *lameContainer) Metrics() (garden.Metrics, error) {
	return garden.Metrics{}, nil
}

func (c *lameContainer) Properties() (garden.Properties, error) {
	return garden.Properties{}, nil
}

func (c *lameContainer) Property(name string) (string, error) {
	return "", nil
}

func (c *lameContainer) SetProperty(name string, value string) error {
	return nil
}

func (c *lameContainer) RemoveProperty(name string) error {
	return nil
}
