// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import swarm "github.com/docker/docker/api/types/swarm"
import types "github.com/docker/docker/api/types"

// NodeAPIClient is an autogenerated mock type for the NodeAPIClient type
type NodeAPIClient struct {
	mock.Mock
}

// NodeInspectWithRaw provides a mock function with given fields: ctx, nodeID
func (_m *NodeAPIClient) NodeInspectWithRaw(ctx context.Context, nodeID string) (swarm.Node, []byte, error) {
	ret := _m.Called(ctx, nodeID)

	var r0 swarm.Node
	if rf, ok := ret.Get(0).(func(context.Context, string) swarm.Node); ok {
		r0 = rf(ctx, nodeID)
	} else {
		r0 = ret.Get(0).(swarm.Node)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, nodeID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, nodeID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NodeList provides a mock function with given fields: ctx, options
func (_m *NodeAPIClient) NodeList(ctx context.Context, options types.NodeListOptions) ([]swarm.Node, error) {
	ret := _m.Called(ctx, options)

	var r0 []swarm.Node
	if rf, ok := ret.Get(0).(func(context.Context, types.NodeListOptions) []swarm.Node); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]swarm.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.NodeListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeRemove provides a mock function with given fields: ctx, nodeID, options
func (_m *NodeAPIClient) NodeRemove(ctx context.Context, nodeID string, options types.NodeRemoveOptions) error {
	ret := _m.Called(ctx, nodeID, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, types.NodeRemoveOptions) error); ok {
		r0 = rf(ctx, nodeID, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NodeUpdate provides a mock function with given fields: ctx, nodeID, version, node
func (_m *NodeAPIClient) NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error {
	ret := _m.Called(ctx, nodeID, version, node)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, swarm.Version, swarm.NodeSpec) error); ok {
		r0 = rf(ctx, nodeID, version, node)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
