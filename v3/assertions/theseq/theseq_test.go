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

package theseq_test

import (
	"slices"
	"testing"

	"github.com/szabba/assert/v3"
	"github.com/szabba/assert/v3/assertions/assertiontesting"
	"github.com/szabba/assert/v3/assertions/theseq"
)

func TestEmpty0(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		// given
		var errfunc assertiontesting.ErrFunc
		var seq theseq.Seq0

		// when
		assert.Using(errfunc.Record).
			That(theseq.Empty0(seq))

		// then
		assert.FailingTest(t).
			That(errfunc.Called()).
			That(errfunc.MessageFormatsTo(`nil sequence (not iterable)`))
	})

	t.Run("NoIterations", func(t *testing.T) {
		// given
		var errfunc assertiontesting.ErrFunc
		var seq theseq.Seq0 = func(_ func() bool) {}

		// when
		assert.Using(errfunc.Record).
			That(theseq.Empty0(seq))

		// then
		assert.FailingTest(t).
			That(errfunc.NotCalled())
	})

	t.Run("SomeIterations", func(t *testing.T) {
		// given
		var errfunc assertiontesting.ErrFunc
		var seq theseq.Seq0 = func(yield func() bool) { yield() }

		// when
		assert.Using(errfunc.Record).
			That(theseq.Empty0(seq))

		// then
		assert.FailingTest(t).
			That(errfunc.Called()).
			That(errfunc.MessageFormatsTo(`sequence is not empty: got 1 iterations, not 0`))
	})
}

func TestEmpty(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		// given
		var errfunc assertiontesting.ErrFunc
		var seq theseq.Seq[int]

		// when
		assert.Using(errfunc.Record).
			That(theseq.Empty(seq))

		// then
		assert.FailingTest(t).
			That(errfunc.Called()).
			That(errfunc.MessageFormatsTo(`nil sequence (not iterable)`))
	})

	t.Run("NoIterations", func(t *testing.T) {
		// given
		var errfunc assertiontesting.ErrFunc
		var seq theseq.Seq[int] = func(_ func(int) bool) {}

		// when
		assert.Using(errfunc.Record).
			That(theseq.Empty(seq))

		// then
		assert.FailingTest(t).
			That(errfunc.NotCalled())
	})

	t.Run("SomeIterations", func(t *testing.T) {
		// given
		var errfunc assertiontesting.ErrFunc
		var seq theseq.Seq[int] = slices.Values([]int{1, 1, 2})

		// when
		assert.Using(errfunc.Record).
			That(theseq.Empty(seq))

		// then
		assert.FailingTest(t).
			That(errfunc.Called()).
			That(errfunc.MessageFormatsTo(`sequence is not empty: got 3 iterations, not 0`))
	})
}

func TestEmpty2(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		// given
		var errfunc assertiontesting.ErrFunc
		var seq theseq.Seq2[int, bool]

		// when
		assert.Using(errfunc.Record).
			That(theseq.Empty2(seq))

		// then
		assert.FailingTest(t).
			That(errfunc.Called()).
			That(errfunc.MessageFormatsTo(`nil sequence (not iterable)`))
	})

	t.Run("NoIterations", func(t *testing.T) {
		// given
		var errfunc assertiontesting.ErrFunc
		var seq theseq.Seq2[int, bool] = func(_ func(int, bool) bool) {}

		// when
		assert.Using(errfunc.Record).
			That(theseq.Empty2(seq))

		// then
		assert.FailingTest(t).
			That(errfunc.NotCalled())
	})

	t.Run("SomeIterations", func(t *testing.T) {
		// given
		var errfunc assertiontesting.ErrFunc
		var seq theseq.Seq2[int, bool] = slices.All([]bool{false, true, false, true})

		// when
		assert.Using(errfunc.Record).
			That(theseq.Empty2(seq))

		// then
		assert.FailingTest(t).
			That(errfunc.Called()).
			That(errfunc.MessageFormatsTo(`sequence is not empty: got 4 iterations, not 0`))
	})
}
