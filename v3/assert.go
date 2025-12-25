// MIT License
//
// Copyright (c) 2022-2025 Karol Marcjan
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

package assert

import "fmt"

// UsingPanic creates an Asserter that panics to report failures.
func UsingPanic() Asserter {
	return Using(nil)
}

// Using creates an Asserter that uses onErr to report failures.
// It is the most general of the Using* functions.
func Using(onErr func(error)) Asserter {
	return Asserter{onErr: onErr}
}

// UsingFmt creates an Asserter that uses fmtFunc to report failures.
func UsingFmt(fmtFunc func(string, ...any)) Asserter {
	onErr := func(err error) { fmtFunc(err.Error()) }
	return Asserter{onErr: onErr}
}

// FailingTest creates an Asserter that fails the test t.Name() and reports all the failed assertions.
func FailingTest(t T) Asserter {
	return Asserter{
		test: t,
		onErr: func(err error) {
			t.Helper()
			t.Error(err)
		},
	}
}

// FailingTest creates an Asserter that fails and interrupts the test t.Name() as soon as the first assertion fails.
func FailingTestFast(t T) Asserter {
	return Asserter{
		test: t,
		onErr: func(err error) {
			t.Helper()
			t.Fatal(err)
		},
	}
}

// An Asserter is used to make assertions.
type Asserter struct {
	test  T
	onErr func(error)
}

// A T contains the subset of the [*testing.T] methods used by Asserters that fail Go tests.
type T interface {
	Helper()
	Error(args ...any)
	Fatal(args ...any)
}

// That asserts there is no problem (ie, the error is nil).
//
// The error func of the asserter receives any non-nil error passed in.
// If the asserter has a nil error func, That panics with the message of the error.
//
// When the assertion passes, the same asserter is returned.
// This enables chaining multiple assertions that share an error func.
//
// When the assertion fails, the error func is called before returning.
// If it panics, the chain is interrupted.
func (a Asserter) That(err error) Asserter {
	if err != nil {
		if a.test != nil {
			a.test.Helper()
		}
		a.fail(err)
	}
	return a
}

// True asserts cond is true.
//
// The error func of the asserter receives msgFmt and args as input.
// If the asserter has a nil error func, True panics with a message formatted by fmt.Sprintf.
//
// When the assertion passes, the same asserter is returned.
// This enables chaining multiple assertions that share and error func.
//
// When the assertion fails, the error func is called before returning.
// If it panics, the chain is interrupted.
func (a Asserter) True(cond bool, msgFmt string, args ...any) Asserter {
	if !cond {
		if a.test != nil {
			a.test.Helper()
		}
		err := fmt.Errorf(msgFmt, args...)
		a.fail(err)
	}
	return a
}

func (a Asserter) fail(err error) {
	if a.test != nil {
		a.test.Helper()
	}
	if a.onErr == nil {
		panic(err.Error())
	}
	a.onErr(err)
}
