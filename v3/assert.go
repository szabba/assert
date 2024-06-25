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

package assert

import "fmt"

// UsingPanic creates an Asserter that panics to report failures.
//
// If you're using this library in tests you probably want UsingFmt instead.
func UsingPanic() Asserter {
	return Using(nil)
}

// Using creates an Asserter that uses onErr to report failures.
//
// Using is the most generic of the Using* functions.
//
// If you're using this library in tests you probably want UsingFmt instead.
func Using(onErr func(error)) Asserter {
	return Asserter{onErr}
}

// UsingFmt creates an Asserter that uses fmtFunc to report failures.
//
// If you're using this library in tests you probably want to call either
//
//	assert.UsingFmt(t.Errorf).That(somethingHolds())
//
// or
//
//	assert.UsingFmt(t.Fatalf).That(somethingHolds())
//
// depending on whether you want the test to continue on failure or not.
func UsingFmt(fmtFunc func(string, ...any)) Asserter {
	onErr := func(err error) { fmtFunc(err.Error()) }
	return Asserter{onErr}
}

// An Asserter is used to make assertions.
type Asserter struct{ onErr func(error) }

// That asserts there is no problem (ie, the error is nil).
//
// The error func of the asserter receives any non-nil error passed in.
// If the asserter has a nil error func, That panics with the message of the error.
//
// When the assertion passes, the same asserter is returned.
// This enables chaining multiple assertions that share an error func.
func (a Asserter) That(err error) Asserter {
	if err != nil {
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
func (a Asserter) True(cond bool, msgFmt string, args ...any) Asserter {
	if !cond {
		err := fmt.Errorf(msgFmt, args...)
		a.fail(err)
	}
	return a
}

func (a Asserter) fail(err error) {
	if a.onErr == nil {
		panic(err.Error())
	}

	a.onErr(err)
}
