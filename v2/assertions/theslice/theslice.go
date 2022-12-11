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

// Package theslice provides reusable assertions about slices.
package theslice

import (
	"fmt"
	"strings"
)

// Empty asserts that s is an empty slice.
func Empty[S ~[]T, T any](s S) (bool, string) {
	if len(s) == 0 {
		return true, ""
	}
	return false, fmt.Sprintf("got non-empty slice %#v", s)
}

// NotEmpty asserts that s is not an empty slice.
func NotEmpty[S ~[]T, T any](s S) (bool, string) {
	if len(s) > 0 {
		return true, ""
	}
	return false, fmt.Sprintf("got empty slice %#v", s)
}

// Equal asserts that an actual slice is equal to an expected one.
//
// Nil slices are never equal to non-nil slices.
// Only slices of equal length can be equal.
// The elements at each index must be equal in both slices.
func Equal[S ~[]T, T comparable](got, want S) (bool, string) {
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

// NotEqual asserts that the actual slice is not equal to another.
//
// For more details look at Equal.
func NotEqual[S ~[]T, T comparable](got, wantNot S) (bool, string) {
	if got == nil && wantNot != nil {
		return true, ""
	}

	if got != nil && wantNot == nil {
		return true, ""
	}

	if len(got) != len(wantNot) {
		return true, ""
	}

	for i := range got {
		if got[i] != wantNot[i] {
			return true, ""
		}
	}

	return false, fmt.Sprintf("got unwanted value %#v", got)
}

// Length asserts the len(s) is n.
func Length[S ~[]T, T any](s S, n int) (bool, string) {
	if len(s) == n {
		return true, ""
	}
	return false, fmt.Sprintf("got slice of length %d, not %d", len(s), n)
}

// LengthNot asserts that len(s) is not n.
func LengthNot[S ~[]T, T any](s S, n int) (bool, string) {
	if len(s) != n {
		return true, ""
	}
	return false, fmt.Sprintf("got slice of length %d", len(s))
}
