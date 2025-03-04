// Code generated by mockery. DO NOT EDIT.

package sym_encrypt

import mock "github.com/stretchr/testify/mock"

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

type MockInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface) EXPECT() *MockInterface_Expecter {
	return &MockInterface_Expecter{mock: &_m.Mock}
}

// Decrypt provides a mock function with given fields: ciphertext
func (_m *MockInterface) Decrypt(ciphertext string) (string, error) {
	ret := _m.Called(ciphertext)

	if len(ret) == 0 {
		panic("no return value specified for Decrypt")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(ciphertext)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(ciphertext)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ciphertext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_Decrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decrypt'
type MockInterface_Decrypt_Call struct {
	*mock.Call
}

// Decrypt is a helper method to define mock.On call
//   - ciphertext string
func (_e *MockInterface_Expecter) Decrypt(ciphertext interface{}) *MockInterface_Decrypt_Call {
	return &MockInterface_Decrypt_Call{Call: _e.mock.On("Decrypt", ciphertext)}
}

func (_c *MockInterface_Decrypt_Call) Run(run func(ciphertext string)) *MockInterface_Decrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockInterface_Decrypt_Call) Return(_a0 string, _a1 error) *MockInterface_Decrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_Decrypt_Call) RunAndReturn(run func(string) (string, error)) *MockInterface_Decrypt_Call {
	_c.Call.Return(run)
	return _c
}

// Encrypt provides a mock function with given fields: plaintext
func (_m *MockInterface) Encrypt(plaintext string) (string, error) {
	ret := _m.Called(plaintext)

	if len(ret) == 0 {
		panic("no return value specified for Encrypt")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(plaintext)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(plaintext)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(plaintext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInterface_Encrypt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encrypt'
type MockInterface_Encrypt_Call struct {
	*mock.Call
}

// Encrypt is a helper method to define mock.On call
//   - plaintext string
func (_e *MockInterface_Expecter) Encrypt(plaintext interface{}) *MockInterface_Encrypt_Call {
	return &MockInterface_Encrypt_Call{Call: _e.mock.On("Encrypt", plaintext)}
}

func (_c *MockInterface_Encrypt_Call) Run(run func(plaintext string)) *MockInterface_Encrypt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockInterface_Encrypt_Call) Return(_a0 string, _a1 error) *MockInterface_Encrypt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInterface_Encrypt_Call) RunAndReturn(run func(string) (string, error)) *MockInterface_Encrypt_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
