// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"

// ExecutableSubWorkflow is an autogenerated mock type for the ExecutableSubWorkflow type
type ExecutableSubWorkflow struct {
	mock.Mock
}

// FromNode provides a mock function with given fields: name
func (_m *ExecutableSubWorkflow) FromNode(name string) ([]string, error) {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnections provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetConnections() *v1alpha1.Connections {
	ret := _m.Called()

	var r0 *v1alpha1.Connections
	if rf, ok := ret.Get(0).(func() *v1alpha1.Connections); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Connections)
		}
	}

	return r0
}

// GetID provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetNode provides a mock function with given fields: nodeID
func (_m *ExecutableSubWorkflow) GetNode(nodeID string) (v1alpha1.ExecutableNode, bool) {
	ret := _m.Called(nodeID)

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableNode); ok {
		r0 = rf(nodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(nodeID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetNodes provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetNodes() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetOnFailureNode provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetOnFailureNode() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}

// GetOutputBindings provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetOutputBindings() []*v1alpha1.Binding {
	ret := _m.Called()

	var r0 []*v1alpha1.Binding
	if rf, ok := ret.Get(0).(func() []*v1alpha1.Binding); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.Binding)
		}
	}

	return r0
}

// GetOutputs provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) GetOutputs() *v1alpha1.OutputVarMap {
	ret := _m.Called()

	var r0 *v1alpha1.OutputVarMap
	if rf, ok := ret.Get(0).(func() *v1alpha1.OutputVarMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.OutputVarMap)
		}
	}

	return r0
}

// StartNode provides a mock function with given fields:
func (_m *ExecutableSubWorkflow) StartNode() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}
