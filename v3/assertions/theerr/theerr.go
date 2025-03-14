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

// Package theerr provides reusable assertions about a single error.
package theerr

import (
	"errors"
	"fmt"
)

// IsNil asserts that the error is nil.
func IsNil(err error) error {
	if err != nil {
		return fmt.Errorf("unexpected error: %w", err)
	}
	return nil
}

// IsNotNil asserts that the error is not nil.
//
// This is less precise than [Is] or [IsA].
// It is better than nothing when there aren't exposed sentinels / types to check for.
func IsNotNil(err error) error {
	if err == nil {
		return errUnexpectedNil
	}
	return nil
}

var errUnexpectedNil = errors.New("unexpected nil error")

// Is asserts that got is the wanted error.
//
// This is aware of error composition.
// For more details see [errors.Is].
func Is(got, want error) error {

	if want == nil {
		return IsNil(got)
	}

	if !errors.Is(got, want) {
		return fmt.Errorf("got error %q, not %q", got, want)
	}

	return nil
}

// IsA asserts that the error can be viewed as an error of type T.
//
// This is aware of error composition.
// For more details see [errors.As].
func IsA[T error](err error) error {

	var typedErr T

	if err == nil {
		return fmt.Errorf("got nil error, not a %T", typedErr)
	}

	if !errors.As(err, &typedErr) {
		return fmt.Errorf("got error of type %T, not %T", err, typedErr)
	}

	return nil
}
