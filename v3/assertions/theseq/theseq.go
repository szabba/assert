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

// Package theseq provides reusable assertions about iterable functions.
package theseq

import "fmt"

type Seq0 = func(yield func() bool)
type Seq[A any] = func(yield func(A) bool)
type Seq2[A, B any] = func(yield func(A, B) bool)

// Empty0 asserts that seq is an empty sequence.
//
// See [Empty] and [Empty2] if your sequence produces values.
func Empty0[S ~Seq0](seq S) error {
	if seq == nil {
		return fmt.Errorf("nil sequence (not iterable)")
	}
	i := 0
	for range seq {
		i++
	}
	if i > 0 {
		return fmt.Errorf("sequence is not empty: got %d iterations, not 0", i)
	}
	return nil
}

// Empty asserts that seq is an empty sequence.
//
// See [Empty0] if your sequence produces no values.
// See [Empty2] if your sequence produces two values each iteration.
func Empty[A any, S ~Seq[A]](seq S) error {
	if seq == nil {
		return fmt.Errorf("nil sequence (not iterable)")
	}
	i := 0
	for range seq {
		i++
	}
	if i > 0 {
		return fmt.Errorf("sequence is not empty: got %d iterations, not 0", i)
	}
	return nil
}

// Empty2 asserts that seq is an empty sequence.
//
// See [Empty0] if your sequence produces no values.
// See [Empty] if your sequence produces one values each iteration.
func Empty2[A, B any, S ~Seq2[A, B]](seq S) error {
	if seq == nil {
		return fmt.Errorf("nil sequence (not iterable)")
	}
	i := 0
	for range seq {
		i++
	}
	if i > 0 {
		return fmt.Errorf("sequence is not empty: got %d iterations, not 0", i)
	}
	return nil
}
