// Code generated by mockery v2.32.2. DO NOT EDIT.

package mocks

import (
	model "github.com/ergomake/layerform/internal/data/model"
	mock "github.com/stretchr/testify/mock"
)

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

type Backend_Expecter struct {
	mock *mock.Mock
}

func (_m *Backend) EXPECT() *Backend_Expecter {
	return &Backend_Expecter{mock: &_m.Mock}
}

// GetLayerState provides a mock function with given fields: layer, instance
func (_m *Backend) GetLayerState(layer *model.Layer, instance string) ([]byte, error) {
	ret := _m.Called(layer, instance)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Layer, string) ([]byte, error)); ok {
		return rf(layer, instance)
	}
	if rf, ok := ret.Get(0).(func(*model.Layer, string) []byte); ok {
		r0 = rf(layer, instance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Layer, string) error); ok {
		r1 = rf(layer, instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Backend_GetLayerState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLayerState'
type Backend_GetLayerState_Call struct {
	*mock.Call
}

// GetLayerState is a helper method to define mock.On call
//   - layer *model.Layer
//   - instance string
func (_e *Backend_Expecter) GetLayerState(layer interface{}, instance interface{}) *Backend_GetLayerState_Call {
	return &Backend_GetLayerState_Call{Call: _e.mock.On("GetLayerState", layer, instance)}
}

func (_c *Backend_GetLayerState_Call) Run(run func(layer *model.Layer, instance string)) *Backend_GetLayerState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Layer), args[1].(string))
	})
	return _c
}

func (_c *Backend_GetLayerState_Call) Return(_a0 []byte, _a1 error) *Backend_GetLayerState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Backend_GetLayerState_Call) RunAndReturn(run func(*model.Layer, string) ([]byte, error)) *Backend_GetLayerState_Call {
	_c.Call.Return(run)
	return _c
}

// SaveLayerState provides a mock function with given fields: layer, instance, _a2
func (_m *Backend) SaveLayerState(layer *model.Layer, instance string, _a2 []byte) error {
	ret := _m.Called(layer, instance, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Layer, string, []byte) error); ok {
		r0 = rf(layer, instance, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Backend_SaveLayerState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveLayerState'
type Backend_SaveLayerState_Call struct {
	*mock.Call
}

// SaveLayerState is a helper method to define mock.On call
//   - layer *model.Layer
//   - instance string
//   - _a2 []byte
func (_e *Backend_Expecter) SaveLayerState(layer interface{}, instance interface{}, _a2 interface{}) *Backend_SaveLayerState_Call {
	return &Backend_SaveLayerState_Call{Call: _e.mock.On("SaveLayerState", layer, instance, _a2)}
}

func (_c *Backend_SaveLayerState_Call) Run(run func(layer *model.Layer, instance string, _a2 []byte)) *Backend_SaveLayerState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Layer), args[1].(string), args[2].([]byte))
	})
	return _c
}

func (_c *Backend_SaveLayerState_Call) Return(_a0 error) *Backend_SaveLayerState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Backend_SaveLayerState_Call) RunAndReturn(run func(*model.Layer, string, []byte) error) *Backend_SaveLayerState_Call {
	_c.Call.Return(run)
	return _c
}

// NewBackend creates a new instance of Backend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBackend(t interface {
	mock.TestingT
	Cleanup(func())
}) *Backend {
	mock := &Backend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}