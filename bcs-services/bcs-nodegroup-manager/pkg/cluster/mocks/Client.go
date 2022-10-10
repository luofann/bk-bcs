// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	cluster "github.com/Tencent/bk-bcs/bcs-services/bcs-nodegroup-manager/pkg/cluster"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// ListClusterNodes provides a mock function with given fields: clusterID
func (_m *Client) ListClusterNodes(clusterID string) ([]*cluster.Node, error) {
	ret := _m.Called(clusterID)

	var r0 []*cluster.Node
	if rf, ok := ret.Get(0).(func(string) []*cluster.Node); ok {
		r0 = rf(clusterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*cluster.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNodesByLabel provides a mock function with given fields: clusterID, labels
func (_m *Client) ListNodesByLabel(clusterID string, labels map[string]interface{}) (map[string]*cluster.Node, error) {
	ret := _m.Called(clusterID, labels)

	var r0 map[string]*cluster.Node
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) map[string]*cluster.Node); ok {
		r0 = rf(clusterID, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*cluster.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]interface{}) error); ok {
		r1 = rf(clusterID, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNodeMetadata provides a mock function with given fields: clusterID, nodeName, labels, annotations
func (_m *Client) UpdateNodeMetadata(clusterID string, nodeName string, labels map[string]interface{}, annotations map[string]interface{}) error {
	ret := _m.Called(clusterID, nodeName, labels, annotations)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, map[string]interface{}, map[string]interface{}) error); ok {
		r0 = rf(clusterID, nodeName, labels, annotations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}