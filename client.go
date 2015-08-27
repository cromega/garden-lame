package main

import (
	"github.com/cloudfoundry-incubator/garden"
)

type lameClient struct {
}

func (c *lameClient) Ping() error {
	return nil
}

func (c *lameClient) Capacity() (garden.Capacity, error) {
	return garden.Capacity{}, nil
}

func (c *lameClient) Create(garden.ContainerSpec) (garden.Container, error) {
	return &lameContainer{}, nil
}

func (c *lameClient) Destroy(handle string) error {
	return nil
}

func (c *lameClient) Containers(garden.Properties) ([]garden.Container, error) {
	return nil, nil
}

func (c *lameClient) BulkInfo(handles []string) (map[string]garden.ContainerInfoEntry, error) {
	return nil, nil
}

func (c *lameClient) BulkMetrics(handles []string) (map[string]garden.ContainerMetricsEntry, error) {
	return nil, nil
}

func (c *lameClient) Lookup(handle string) (garden.Container, error) {
	return &lameContainer{}, nil
}
