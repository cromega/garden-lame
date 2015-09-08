package main

import (
	"github.com/cloudfoundry-incubator/garden"
	"github.com/pivotal-golang/lager"
)

type lameClient struct {
	logger     lager.Logger
	containers map[string]*lameContainer
}

func NewLameClient(logger lager.Logger) *lameClient {
	return &lameClient{logger: logger, containers: make(map[string]*lameContainer)}
}

func (c *lameClient) Ping() error {
	return nil
}

func (c *lameClient) Capacity() (garden.Capacity, error) {
	return garden.Capacity{}, nil
}

func (c *lameClient) Create(spec garden.ContainerSpec) (garden.Container, error) {
	container := NewLameContainer(spec.Handle)
	c.containers[spec.Handle] = container
	return container, nil
}

func (c *lameClient) Destroy(handle string) error {
	return nil
}

func (c *lameClient) Containers(garden.Properties) ([]garden.Container, error) {
	c.logger.Info("containers")
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
