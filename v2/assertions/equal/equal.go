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

// Package equal provides reusable assertions about equal values.
package equal

import (
	"fmt"
	"strings"
)

// Values asserts that two values of a [comparable] type are equal.
//
// [comparable]: https://go.dev/ref/spec#Type_constraints
func Values[T comparable](got, want T) (bool, string) {
	if got == want {
		return true, ""
	}
	return false, fmt.Sprintf("got %#v, not %#v", got, want)
}

// Slices asserts that two slices with element of a [comparable] type are equal.
//
// Slices supports named slice types, ie, if you define MySlice as
//
//	type MySlice []int
//
// you can pass values of type MySlice to Slices.
//
// [comparable]: https://go.dev/ref/spec#Type_constraints
func Slices[T comparable, S ~[]T](got, want S) (bool, string) {
	if got == nil && want != nil {
		msg := fmt.Sprintf("got nil, not %#v", want)
		return false, msg
	}

	if got != nil && want == nil {
		msg := fmt.Sprintf("got %#v, not nil", got)
		return false, msg
	}

	if len(got) != len(want) {
		msg := fmt.Sprintf(
			"got %#v (of length %d), not %#v (of length %d)",
			got, len(got), want, len(want))
		return false, msg
	}

	diffs := make([]int, 0, len(got))
	for i := range got {
		if got[i] != want[i] {
			diffs = append(diffs, i)
		}
	}

	diffMsgs := make([]string, len(diffs))
	for i, d := range diffs {
		diffMsgs[i] = fmt.Sprintf(
			"element at position %d is %#v, not %#v",
			d, got[d], want[d])
	}

	if len(diffs) > 0 {
		msg := fmt.Sprintf(
			"got slice %#v, not %#v: %s",
			got, want, strings.Join(diffMsgs, "; "))
		return false, msg
	}

	return true, ""
}
