// MIT License
//
// Copyright (c) 2022 Karol Marcjan
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package assertiontesting

import (
	"errors"
	"fmt"
)

var (
	errErrFuncNotCalled = errors.New("error function was not called")
	errErrFuncCalled    = errors.New("error function was called")
)

// ErrFunc is an object that records the details of a single call to an error function.
type ErrFunc struct {
	called bool

	err error
}

// Record is a method that can be used as an error function.
//
// Pass f.Record to assert.Using in order to test an assertion helper.
//
// If Record gets called, the details of the call are recorded.
// If Record is called more than once, it panics.
//
// The other methods on the object can be used to make assertions about the call.
// This works both when Record was actually called and not.
func (f *ErrFunc) Record(err error) {

	if f.called {
		panic("error function called multiple times")
	}

	if err == nil {
		panic("error function called with a nil error")
	}

	f.called = true
	f.err = err
}

// Called is a helper to assert that the error function was called.
//
// See Record for more details.
func (f *ErrFunc) Called() error {

	if !f.called {
		return errErrFuncNotCalled
	}
	return nil
}

// NotCalled is a helper to assert that the error func was not called.
//
// See Record for more details.
func (f *ErrFunc) NotCalled() error {

	if f.called {
		return errErrFuncCalled
	}
	return nil
}

// MessageFormatsTo is a helper to assert the formatting of the message the error function was called with.
//
// See Record for more details.
func (f *ErrFunc) MessageFormatsTo(want string) error {

	if !f.called {
		return errErrFuncCalled
	}

	got := f.err.Error()
	if got != want {
		return fmt.Errorf("got assertion error with message %q, not %q", got, want)
	}
	return nil
}
