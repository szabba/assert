// MIT License
//
// Copyright (c) 2018 Karol Marcjan
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

// Using creates an Asserter uses onError to report failures.
func Using(onErr ErrorFunc) Asserter {
	return Asserter{onErr}
}

// An ErrorFunc describes what to do when an assertion fails.
type ErrorFunc func(msgFmt string, args ...any)

// An Asserter is used to make assertions.
//
// For failed assertions, the zero asserter formats the message and panics.
type Asserter struct{ onErr ErrorFunc }

// That asserts cond is true.
//
// The ErrorFunc of the Asserter receives msgFmt and args as input.
// If the Asserter has a nil ErrorFunc, That panics with the formatted message.
//
// The same asserter is returned to enable chaining multiple assertions on the same ErrorFunc.
func (a Asserter) That(cond bool, msgFmt string, args ...any) Asserter {
	if !cond {
		a.fail(msgFmt, args)
	}
	return a
}

func (a Asserter) fail(msgFmt string, args []any) {
	if a.onErr == nil {
		msg := fmt.Sprintf(msgFmt, args...)
		panic(msg)
	}

	a.onErr(msgFmt, args...)
}
