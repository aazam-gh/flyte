// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"

// NodeStatusGetter is an autogenerated mock type for the NodeStatusGetter type
type NodeStatusGetter struct {
	mock.Mock
}

// GetNodeExecutionStatus provides a mock function with given fields: id
func (_m *NodeStatusGetter) GetNodeExecutionStatus(id string) v1alpha1.ExecutableNodeStatus {
	ret := _m.Called(id)

	var r0 v1alpha1.ExecutableNodeStatus
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableNodeStatus); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNodeStatus)
		}
	}

	return r0
}
