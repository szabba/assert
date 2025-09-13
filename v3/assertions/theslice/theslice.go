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

// Package theslice provides reusable assertions about slices.
package theslice

import (
	"errors"
	"fmt"
)

// Empty asserts that s is an empty slice.
func Empty[S ~[]T, T any](s S) error {
	if len(s) != 0 {
		return fmt.Errorf("got non-empty slice %#v", s)
	}
	return nil
}

// NotEmpty asserts that s is not an empty slice.
func NotEmpty[S ~[]T, T any](s S) error {
	if len(s) == 0 {
		return fmt.Errorf("got empty slice %#v", s)
	}
	return nil
}

// Equal asserts that an actual slice is equal to an expected one, element-by-element.
//
// Nil slices are never equal to non-nil slices.
// Only slices of equal length can be equal.
// The elements at each index must be equal in both slices.
func Equal[S ~[]T, T comparable](got, want S) error {
	return EqualFunc(got, want, func(l, r T) bool { return l == r })
}

// EqualFunc asserts that an actual slice is equal to an expected one, element-by-element.
// Elements are compared for equality using the function eq.
//
// Nil slices are never equal to non-nil slices.
// Only slices of equal length can be equal.
// The elements at each index must be equal in both slices, as determined by eq.
func EqualFunc[S ~[]T, T any](got, want S, eq func(T, T) bool) error {

	if got == nil && want != nil {
		return fmt.Errorf("got nil, not %#v", want)
	}

	if got != nil && want == nil {
		return fmt.Errorf("got %#v, not nil", got)
	}

	return EqualElementsFunc(got, want, eq)
}

// EqualElements asserts that two slices have all elements equal, element-by-element.
//
// Nil and empty slices compare equal.
// Only slices of equal length can have equal.
// The elements at each index must be equal in both slices.
func EqualElements[S ~[]T, T comparable](got, want S) error {
	return EqualElementsFunc(got, want, func(l, r T) bool { return l == r })
}

// EqualElementsFunc asserts that two slices have all elements equal, element-by-element.
// Elements are compared for equality using the function eq.
//
// Nil and empty slices compare equal.
// Only slices of equal length can have equal.
// The elements at each index must be equal in both slices, as determined by eq.
func EqualElementsFunc[S ~[]T, T any](got, want S, eq func(T, T) bool) error {

	if len(got) != len(want) {
		return fmt.Errorf(
			"got %#v (of length %d), not %#v (of length %d)",
			got, len(got), want, len(want))
	}

	diffs := make([]int, 0, len(got))
	for i := range got {
		if !eq(got[i], want[i]) {
			diffs = append(diffs, i)
		}
	}

	errs := make([]error, 0, len(diffs))
	for _, d := range diffs {
		errs = append(errs, fmt.Errorf(
			"element at position %d is %#v, not %#v",
			d, got[d], want[d]))
	}

	if len(diffs) > 0 {
		sep := " "
		if len(errs) > 1 {
			sep = "\n"
		}
		return fmt.Errorf(
			"got slice %#v, not %#v:%s%w",
			got, want, sep, errors.Join(errs...))
	}

	return nil
}

// NotEqual asserts that the actual slice is not equal to another, element-by-element.
//
// For more details look at Equal.
func NotEqual[S ~[]T, T comparable](got, wantNot S) error {
	if got == nil && wantNot != nil {
		return nil
	}

	if got != nil && wantNot == nil {
		return nil
	}

	if len(got) != len(wantNot) {
		return nil
	}

	for i := range got {
		if got[i] != wantNot[i] {
			return nil
		}
	}

	return fmt.Errorf("got unwanted value %#v", got)
}

// Length asserts the len(s) is n.
func Length[S ~[]T, T any](s S, n int) error {
	if len(s) != n {
		return fmt.Errorf("got slice of length %d, not %d", len(s), n)
	}
	return nil
}

// LengthNot asserts that len(s) is not n.
func LengthNot[S ~[]T, T any](s S, n int) error {
	if len(s) == n {
		return fmt.Errorf("got slice of length %d", len(s))
	}
	return nil
}

// LengthAtLeast asserts that len(s) >= n.
func LengthAtLeast[S ~[]T, T any](s S, n int) error {
	if len(s) < n {
		return fmt.Errorf("got slice of length %d, less than %d", len(s), n)
	}
	return nil
}

// At assers that the i-th element of s passes the assertion in f.
func At[S ~[]T, T any](s S, i int, f func(t T) error) error {
	err := LengthAtLeast(s, i+1)
	if err != nil {
		return err
	}

	el := s[i]
	err = f(el)
	if err != nil {
		return fmt.Errorf("element %d, %#v does not pass: %w", i, el, err)
	}

	return nil
}
