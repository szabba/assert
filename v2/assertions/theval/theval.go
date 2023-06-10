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

// Package theval provides the most basic reusable assertions.
package theval

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

// Equal asserts that an actual value is equal to an expected value of the same.
func Equal[T comparable](got, want T) (bool, string) {
	if got == want {
		return true, ""
	}
	return false, fmt.Sprintf("got %#v, not %#v", got, want)
}

// NotEqual asserts that an actual value is not equal to another.
func NotEqual[T comparable](got, wantNot T) (bool, string) {
	if got != wantNot {
		return true, ""
	}
	return false, fmt.Sprintf("got unwanted value %#v", got)
}

// LessThan asserts that an actual value is less than another.
func LessThan[T constraints.Ordered](got, want T) (bool, string) {
	if got < want {
		return true, ""
	}

	msg := fmt.Sprintf("got %v >= %v", got, want)
	return false, msg
}

// Zero asserts that v is the zero value of it's underlying type.
func Zero(v any) (bool, string) {
	rv := reflect.ValueOf(v)

	if rv.IsZero() {
		return true, ""
	}

	z := reflect.Zero(rv.Type())
	return false, fmt.Sprintf("got %#v, not zero value %#v", v, z)
}

// NotZero asserts that v is not the zero value of it's underlying type.
func NotZero(v any) (bool, string) {
	if !reflect.ValueOf(v).IsZero() {
		return true, ""
	}
	return false, fmt.Sprintf("got zero value %#v", v)
}
