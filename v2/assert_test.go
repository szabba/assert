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

package assert_test

import (
	"testing"

	"github.com/szabba/assert/v2"
)

func TestPassingAssertionWithNilErrorFuncDoesNotPanic(t *testing.T) {
	// given
	// when
	p := catchPanic(func() { assert.UsingPanic().That(true, "OK") })

	// then
	if p != nil {
		t.Errorf("unexpected panic: %#v", p)
	}
}

func TestFailingAssertionWithNilErrorFuncPanicsWithTheFormattedMessage(t *testing.T) {
	// given
	// when
	p := catchPanic(func() { assert.UsingPanic().That(false, "Oops: %#v", false) })

	// then
	wantMsg := "Oops: false"
	if p != wantMsg {
		t.Errorf("got panic %#v, not %q", p, wantMsg)
	}
}

func TestPassingAssertionDoesNotCallNonNilErrorFunc(t *testing.T) {
	// given
	called := false
	errFunc := func(_ string, _ ...any) { called = true }

	// when
	assert.Using(errFunc).That(true, "OK")

	// then
	if called {
		t.Error("the ErrorFunc was called")
	}
}

func TestFailingAssertionCallsNonNilErrorFunc(t *testing.T) {
	// given
	var got struct {
		called bool
		msgFmt string
		args   []any
	}
	errFunc := func(msgFmt string, args ...any) {
		got.called = true
		got.msgFmt, got.args = msgFmt, args
	}

	// when
	assert.Using(errFunc).That(false, "Oops: %#v", false)

	// then
	if !got.called {
		t.Errorf("ErrorFunc was not called")
	}

	wantMsgFmt, wantArgs := "Oops %#v", []any{false}

	if got.msgFmt != "Oops: %#v" {
		t.Errorf("ErrorFunc got msgFmt %q, not %q", got.msgFmt, wantMsgFmt)
	}

	if len(got.args) != 1 || got.args[0] != false {
		t.Errorf("ErrorFunc got args %v, not %v", got.args, wantArgs)
	}
}

func catchPanic(f func()) (caught any) {
	defer func() { caught = recover() }()
	f()
	return caught
}
