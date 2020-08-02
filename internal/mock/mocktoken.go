package mock

import "time"

type MockToken struct {
	MockWait        func() bool
	MockWaitTimeout func(time.Duration) bool
	MockError       func() error
}

func (t *MockToken) Wait() bool {
	return t.MockWait()
}

func (t *MockToken) WaitTimeout(d time.Duration) bool {
	return t.MockWaitTimeout(d)
}

func (t *MockToken) Error() error {
	return t.Error()
}
