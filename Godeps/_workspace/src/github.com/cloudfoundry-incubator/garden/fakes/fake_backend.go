// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/garden"
)

type FakeBackend struct {
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 error
	}
	CapacityStub        func() (garden.Capacity, error)
	capacityMutex       sync.RWMutex
	capacityArgsForCall []struct{}
	capacityReturns     struct {
		result1 garden.Capacity
		result2 error
	}
	CreateStub        func(garden.ContainerSpec) (garden.Container, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 garden.ContainerSpec
	}
	createReturns struct {
		result1 garden.Container
		result2 error
	}
	DestroyStub        func(handle string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		handle string
	}
	destroyReturns struct {
		result1 error
	}
	ContainersStub        func(garden.Properties) ([]garden.Container, error)
	containersMutex       sync.RWMutex
	containersArgsForCall []struct {
		arg1 garden.Properties
	}
	containersReturns struct {
		result1 []garden.Container
		result2 error
	}
	BulkInfoStub        func(handles []string) (map[string]garden.ContainerInfoEntry, error)
	bulkInfoMutex       sync.RWMutex
	bulkInfoArgsForCall []struct {
		handles []string
	}
	bulkInfoReturns struct {
		result1 map[string]garden.ContainerInfoEntry
		result2 error
	}
	BulkMetricsStub        func(handles []string) (map[string]garden.ContainerMetricsEntry, error)
	bulkMetricsMutex       sync.RWMutex
	bulkMetricsArgsForCall []struct {
		handles []string
	}
	bulkMetricsReturns struct {
		result1 map[string]garden.ContainerMetricsEntry
		result2 error
	}
	LookupStub        func(handle string) (garden.Container, error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		handle string
	}
	lookupReturns struct {
		result1 garden.Container
		result2 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	StopStub             func()
	stopMutex            sync.RWMutex
	stopArgsForCall      []struct{}
	GraceTimeStub        func(garden.Container) time.Duration
	graceTimeMutex       sync.RWMutex
	graceTimeArgsForCall []struct {
		arg1 garden.Container
	}
	graceTimeReturns struct {
		result1 time.Duration
	}
}

func (fake *FakeBackend) Ping() error {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	} else {
		return fake.pingReturns.result1
	}
}

func (fake *FakeBackend) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeBackend) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackend) Capacity() (garden.Capacity, error) {
	fake.capacityMutex.Lock()
	fake.capacityArgsForCall = append(fake.capacityArgsForCall, struct{}{})
	fake.capacityMutex.Unlock()
	if fake.CapacityStub != nil {
		return fake.CapacityStub()
	} else {
		return fake.capacityReturns.result1, fake.capacityReturns.result2
	}
}

func (fake *FakeBackend) CapacityCallCount() int {
	fake.capacityMutex.RLock()
	defer fake.capacityMutex.RUnlock()
	return len(fake.capacityArgsForCall)
}

func (fake *FakeBackend) CapacityReturns(result1 garden.Capacity, result2 error) {
	fake.CapacityStub = nil
	fake.capacityReturns = struct {
		result1 garden.Capacity
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) Create(arg1 garden.ContainerSpec) (garden.Container, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 garden.ContainerSpec
	}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeBackend) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeBackend) CreateArgsForCall(i int) garden.ContainerSpec {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *FakeBackend) CreateReturns(result1 garden.Container, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 garden.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) Destroy(handle string) error {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		handle string
	}{handle})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(handle)
	} else {
		return fake.destroyReturns.result1
	}
}

func (fake *FakeBackend) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeBackend) DestroyArgsForCall(i int) string {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].handle
}

func (fake *FakeBackend) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackend) Containers(arg1 garden.Properties) ([]garden.Container, error) {
	fake.containersMutex.Lock()
	fake.containersArgsForCall = append(fake.containersArgsForCall, struct {
		arg1 garden.Properties
	}{arg1})
	fake.containersMutex.Unlock()
	if fake.ContainersStub != nil {
		return fake.ContainersStub(arg1)
	} else {
		return fake.containersReturns.result1, fake.containersReturns.result2
	}
}

func (fake *FakeBackend) ContainersCallCount() int {
	fake.containersMutex.RLock()
	defer fake.containersMutex.RUnlock()
	return len(fake.containersArgsForCall)
}

func (fake *FakeBackend) ContainersArgsForCall(i int) garden.Properties {
	fake.containersMutex.RLock()
	defer fake.containersMutex.RUnlock()
	return fake.containersArgsForCall[i].arg1
}

func (fake *FakeBackend) ContainersReturns(result1 []garden.Container, result2 error) {
	fake.ContainersStub = nil
	fake.containersReturns = struct {
		result1 []garden.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) BulkInfo(handles []string) (map[string]garden.ContainerInfoEntry, error) {
	fake.bulkInfoMutex.Lock()
	fake.bulkInfoArgsForCall = append(fake.bulkInfoArgsForCall, struct {
		handles []string
	}{handles})
	fake.bulkInfoMutex.Unlock()
	if fake.BulkInfoStub != nil {
		return fake.BulkInfoStub(handles)
	} else {
		return fake.bulkInfoReturns.result1, fake.bulkInfoReturns.result2
	}
}

func (fake *FakeBackend) BulkInfoCallCount() int {
	fake.bulkInfoMutex.RLock()
	defer fake.bulkInfoMutex.RUnlock()
	return len(fake.bulkInfoArgsForCall)
}

func (fake *FakeBackend) BulkInfoArgsForCall(i int) []string {
	fake.bulkInfoMutex.RLock()
	defer fake.bulkInfoMutex.RUnlock()
	return fake.bulkInfoArgsForCall[i].handles
}

func (fake *FakeBackend) BulkInfoReturns(result1 map[string]garden.ContainerInfoEntry, result2 error) {
	fake.BulkInfoStub = nil
	fake.bulkInfoReturns = struct {
		result1 map[string]garden.ContainerInfoEntry
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) BulkMetrics(handles []string) (map[string]garden.ContainerMetricsEntry, error) {
	fake.bulkMetricsMutex.Lock()
	fake.bulkMetricsArgsForCall = append(fake.bulkMetricsArgsForCall, struct {
		handles []string
	}{handles})
	fake.bulkMetricsMutex.Unlock()
	if fake.BulkMetricsStub != nil {
		return fake.BulkMetricsStub(handles)
	} else {
		return fake.bulkMetricsReturns.result1, fake.bulkMetricsReturns.result2
	}
}

func (fake *FakeBackend) BulkMetricsCallCount() int {
	fake.bulkMetricsMutex.RLock()
	defer fake.bulkMetricsMutex.RUnlock()
	return len(fake.bulkMetricsArgsForCall)
}

func (fake *FakeBackend) BulkMetricsArgsForCall(i int) []string {
	fake.bulkMetricsMutex.RLock()
	defer fake.bulkMetricsMutex.RUnlock()
	return fake.bulkMetricsArgsForCall[i].handles
}

func (fake *FakeBackend) BulkMetricsReturns(result1 map[string]garden.ContainerMetricsEntry, result2 error) {
	fake.BulkMetricsStub = nil
	fake.bulkMetricsReturns = struct {
		result1 map[string]garden.ContainerMetricsEntry
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) Lookup(handle string) (garden.Container, error) {
	fake.lookupMutex.Lock()
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		handle string
	}{handle})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(handle)
	} else {
		return fake.lookupReturns.result1, fake.lookupReturns.result2
	}
}

func (fake *FakeBackend) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeBackend) LookupArgsForCall(i int) string {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return fake.lookupArgsForCall[i].handle
}

func (fake *FakeBackend) LookupReturns(result1 garden.Container, result2 error) {
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 garden.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) Start() error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	} else {
		return fake.startReturns.result1
	}
}

func (fake *FakeBackend) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeBackend) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackend) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		fake.StopStub()
	}
}

func (fake *FakeBackend) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeBackend) GraceTime(arg1 garden.Container) time.Duration {
	fake.graceTimeMutex.Lock()
	fake.graceTimeArgsForCall = append(fake.graceTimeArgsForCall, struct {
		arg1 garden.Container
	}{arg1})
	fake.graceTimeMutex.Unlock()
	if fake.GraceTimeStub != nil {
		return fake.GraceTimeStub(arg1)
	} else {
		return fake.graceTimeReturns.result1
	}
}

func (fake *FakeBackend) GraceTimeCallCount() int {
	fake.graceTimeMutex.RLock()
	defer fake.graceTimeMutex.RUnlock()
	return len(fake.graceTimeArgsForCall)
}

func (fake *FakeBackend) GraceTimeArgsForCall(i int) garden.Container {
	fake.graceTimeMutex.RLock()
	defer fake.graceTimeMutex.RUnlock()
	return fake.graceTimeArgsForCall[i].arg1
}

func (fake *FakeBackend) GraceTimeReturns(result1 time.Duration) {
	fake.GraceTimeStub = nil
	fake.graceTimeReturns = struct {
		result1 time.Duration
	}{result1}
}

var _ garden.Backend = new(FakeBackend)
