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
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/szabba/assert/v3"
)

var _ assert.T = new(testing.T)

func TestFailingTest(t *testing.T) {

	t.Run("DoesNothingWithNoError", func(t *testing.T) {
		// given
		fakeT := new(FakeT)

		// when
		assert.FailingTest(fakeT).That(nil)

		// then
		fakeT.LogCalls(t)
		if len(fakeT.Calls) > 0 {
			t.Error("calls were made to the fake *testing.T when no error was present")
		}
	})

	t.Run("CallsHelperThenErrorWhenErrorIsReported", func(t *testing.T) {
		// given
		fakeT := new(FakeT)

		err := io.ErrUnexpectedEOF

		// when
		assert.FailingTest(fakeT).That(err)

		// then
		fakeT.LogCalls(t)

		if len(fakeT.Calls) < 2 {
			t.Fatal("less than two calls to the fake *testing.T were made")
		}

		prefix := fakeT.Calls[:len(fakeT.Calls)-1]
		last := fakeT.Calls[len(fakeT.Calls)-1]

		for i, c := range prefix {
			if c.Method != "Helper" {
				t.Errorf("call number %d was not a call to the fake *testing.T's Helper method", i)
			}
		}

		if last.Method != "Error" {
			t.Error("the last call was not to the fake *testing.T's Error method")
		}

		if len(last.Args) < 1 {
			t.Fatal("the call to the fake *testing.T's Error method got no arguments")
		}

		if len(last.Args) > 1 {
			t.Fatal("the call to the fake *testing.T's Error method got multiple arguments")
		}

		if last.Args[0] != err {
			t.Errorf("reported error %q, not %q", fakeT.Calls[1].Args[0], err)
		}
	})
}

func TestFailingTestFast(t *testing.T) {

	t.Run("DoesNothingWithNoError", func(t *testing.T) {
		// given
		fakeT := new(FakeT)

		// when
		assert.FailingTestFast(fakeT).That(nil)

		// then
		fakeT.LogCalls(t)
		if len(fakeT.Calls) > 0 {
			t.Error("calls were made to the fake *testing.T when no error was present")
		}
	})

	t.Run("CallsHelperThenErrorWhenErrorIsReported", func(t *testing.T) {
		// given
		fakeT := new(FakeT)

		err := io.ErrUnexpectedEOF

		// when
		assert.FailingTestFast(fakeT).That(err)

		// then
		fakeT.LogCalls(t)

		if len(fakeT.Calls) < 2 {
			t.Fatal("less than two calls to the fake *testing.T were made")
		}

		prefix := fakeT.Calls[:len(fakeT.Calls)-1]
		last := fakeT.Calls[len(fakeT.Calls)-1]

		for i, c := range prefix {
			if c.Method != "Helper" {
				t.Errorf("call number %d was not a call to the fake *testing.T's Helper method", i)
			}
		}

		if last.Method != "Fatal" {
			t.Error("the last call was not to the fake *testing.T's Fatal method")
		}

		if len(last.Args) < 1 {
			t.Fatal("the call to the fake *testing.T's Fatal method got no arguments")
		}

		if len(last.Args) > 1 {
			t.Fatal("the call to the fake *testing.T's Fatal method got multiple arguments")
		}

		if last.Args[0] != err {
			t.Errorf("reported error %q, not %q", fakeT.Calls[1].Args[0], err)
		}
	})
}

var _ assert.T = new(FakeT)

type FakeT struct {
	Calls []FakeTCall
}

type FakeTCall struct {
	Method string
	Args   []any // nil when Method == "Helper"
}

func (f *FakeT) Helper() {
	f.Calls = append(f.Calls, FakeTCall{Method: "Helper"})
}

func (f *FakeT) Error(args ...any) {
	call := FakeTCall{
		Method: "Error",
		Args:   args,
	}
	f.Calls = append(f.Calls, call)
}

func (f *FakeT) Fatal(args ...any) {
	call := FakeTCall{
		Method: "Fatal",
		Args:   args,
	}
	f.Calls = append(f.Calls, call)
}

func (f *FakeT) LogCalls(t *testing.T) {
	for _, c := range f.Calls {
		bld := new(strings.Builder)
		bld.WriteString(c.Method)
		bld.WriteRune('(')
		for i, a := range c.Args {
			if i > 0 {
				bld.WriteString(", ")
			}
			fmt.Fprint(bld, a)
		}
		bld.WriteRune(')')
		t.Log(bld.String())
	}
}
