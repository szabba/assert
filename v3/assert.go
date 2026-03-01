// MIT License
//
// Copyright (c) 2022-2026 Karol Marcjan
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
// If the error-reporting function does not panic, the asserter is returned to allow method chaining.
func (a Asserter) That(err error) Asserter {
	if err != nil {
		if a.test != nil {
			a.test.Helper()
		}
		a.fail(err)
	}
	return a
}

// Thatf asserts that there is no problem (ie, the error is nil) providing additional context about it.
//
// If the error-reporting function does not panic, the asserter is returned to allow method chaining.
func (a Asserter) Thatf(err error, msgFmt string, args ...any) Asserter {
	if err != nil {
		if a.test != nil {
			a.test.Helper()
		}
		args = append(append([]any{}, args...), err)
		err = fmt.Errorf(msgFmt+": %w", args...)
		a.fail(err)
	}
	return a
}

// True asserts cond is true.
// When cond is false, msgFmt and args are passed to [fmt.Errorf] to build the reported error.
//
// If the error-reporting function does not panic, the asserter is returned to allow method chaining.
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
