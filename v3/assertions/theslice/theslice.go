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
	return EqualFunc(got, want, comparableEq)
}

// EqualFunc asserts that an actual slice is equal to an expected one, element-by-element.
// Elements are compared for equality using the function eq.
//
// Nil slices are never equal to non-nil slices.
// Only slices of equal length can be equal.
// The elements at each index must be equal in both slices, as determined by eq.
//
// Deprecated: Prefer [EqualErrFunc] - it produces better error messages.
func EqualFunc[S ~[]T, T any](got, want S, eq func(T, T) bool) error {
	return EqualErrFunc(got, want, func(got, want T) error {
		if !eq(got, want) {
			return fmt.Errorf("got %#v, want %#v", got, want)
		}
		return nil
	})
}

// EqualFunc asserts that an actual slice is equal to an expected one, element-by-element.
// Elements are compared for equality using the function eq; a nil error indicates equality.
//
// Nil slices are never equal to non-nil slices.
// Only slices of equal length can be equal.
// The elements at each index must be equal in both slices, as determined by eq.
func EqualErrFunc[S ~[]T, T any](got, want S, eq func(T, T) error) error {
	if got == nil && want != nil {
		return fmt.Errorf("got nil, not %#v", want)
	}

	if got != nil && want == nil {
		return fmt.Errorf("got %#v, not nil", got)
	}

	return EqualElementsErrFunc(got, want, eq)
}

// EqualElements asserts that two slices have all elements equal, element-by-element.
//
// Nil and empty slices compare equal.
// Only slices of equal length can have equal.
// The elements at each index must be equal in both slices.
func EqualElements[S ~[]T, T comparable](got, want S) error {
	return EqualElementsFunc(got, want, comparableEq)
}

// EqualElementsFunc asserts that two slices have all elements equal, element-by-element.
// Elements are compared for equality using the function eq.
//
// Nil and empty slices compare equal.
// Only slices of equal length can have equal.
// The elements at each index must be equal in both slices, as determined by eq.
//
// Deprecated: Prefer [EqualElementsErrFunc] - it produces better error messages.
func EqualElementsFunc[S ~[]T, T any](got, want S, eq func(T, T) bool) error {
	return EqualElementsErrFunc(got, want, func(got, want T) error {
		if !eq(got, want) {
			return fmt.Errorf("got %#v, want %#v", got, want)
		}
		return nil
	})
}

// EqualElementErrFunc asserts that two slices have all elements equal, element-by-element.
// Elements are compared for equality using the function eq; a nil error indicates equality.
//
// Nil and empty slices compare equal.
// Only slices of equal length can have equal.
// The elements at each index must be equal in both slices, as determined by eq.
func EqualElementsErrFunc[S ~[]T, T any](got, want S, eq func(T, T) error) error {

	if len(got) != len(want) {
		return fmt.Errorf(
			"got %#v (of length %d), not %#v (of length %d)",
			got, len(got), want, len(want))
	}

	errs := make([]error, 0, len(got))
	for i := range got {
		err := eq(got[i], want[i])
		if err != nil {
			errs = append(errs, fmt.Errorf(
				"element at position %d is %#v, not %#v",
				i, got[i], want[i]))
		}
	}

	if len(errs) > 0 {
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

// At asserts that the i-th element of s passes the assertion in f.
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

// IsPrefix asserts that gotPrefix is a non-strict prefix of another slice.
//
// The first len(gotPrefix) elements of both slices must be equal.
// Because IsPrefix checks for a non-strict prefix, gotPrefix and of can have equal length.
func IsPrefix[S ~[]T, T comparable](gotPrefix, of S) error {
	return IsPrefixFunc(gotPrefix, of, comparableEq)
}

// IsPrefixFunc asserts that gotPrefix is a non-strict prefix of another slice.
// Elements are compared for equality using the function eq.
//
// The first len(gotPrefix) elements of both slices must be equal.
// Because IsPrefixFunc checks for a non-strict prefix, gotPrefix and of can have equal length.
func IsPrefixFunc[S ~[]T, T any](gotPrefix, of S, eq func(T, T) bool) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("got slice %#v, not a prefix of %#v: %w", gotPrefix, of, err)
		}
	}()

	if len(gotPrefix) > len(of) {
		return fmt.Errorf(
			"actual slice longer (%d) than reference (%d)",
			len(gotPrefix), len(of))
	}

	for i := range gotPrefix {
		if !eq(gotPrefix[i], of[i]) {
			return fmt.Errorf(
				"slices start differing at index %d: got %#v, not %#v",
				i, gotPrefix[i], of[i])
		}
	}

	return nil
}

// HasPrefix asserts that got has another slice as a non-strict prefix.
//
// The first len(wantPrefix) elements of both slices must be equal.
// Because HasPrefix checks for a non-strict prefix, gotPrefix and of can have equal length.
func HasPrefix[S ~[]T, T comparable](got, wantPrefix S) error {
	return HasPrefixFunc(got, wantPrefix, comparableEq)
}

// HasPrefix asserts that got has another slice as a non-strict prefix.
// Elements are compared for equality using the function eq.
//
// The first len(wantPrefix) elements of both slices must be equal.
// Because HasPrefix checks for a non-strict prefix, gotPrefix and of can have equal length.
func HasPrefixFunc[S ~[]T, T any](got, wantPrefix S, eq func(T, T) bool) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("got %#v, not slice with prefix %#v: %w", got, wantPrefix, err)
		}
	}()

	if len(got) < len(wantPrefix) {
		return fmt.Errorf(
			"slice shorter (%d) than expected prefix (%d)",
			len(got), len(wantPrefix))
	}

	for i := range wantPrefix {
		if !eq(got[i], wantPrefix[i]) {
			return fmt.Errorf(
				"slices start differing at index %d: got %#v, not %#v",
				i, got[i], wantPrefix[i])
		}
	}

	return nil
}

func comparableEq[T comparable](l, r T) bool { return l == r }
