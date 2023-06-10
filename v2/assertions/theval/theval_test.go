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

package theval_test

import (
	"testing"

	"github.com/szabba/assert/v2"
	"github.com/szabba/assert/v2/assertions/assertiontesting"

	"github.com/szabba/assert/v2/assertions/theval"
)

func TestEqual(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.Equal(0, 0))

		// then
		assert.
			Using(t.Errorf).
			That(errFunc.NotCalled())
	})

	t.Run("False", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.Equal(0, 1))

		// then
		assert.
			Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got 0, not 1"))
	})

}

func TestNotEqual(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.NotEqual(0, 1))

		// then
		assert.
			Using(t.Errorf).
			That(errFunc.NotCalled())
	})

	t.Run("False", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.NotEqual(0, 0))

		// then
		assert.
			Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got unwanted value 0"))
	})

}

func TestLessThan(t *testing.T) {

	t.Run("1vs3", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.LessThan(1, 3))

		// then
		assert.
			Using(t.Errorf).
			That(errFunc.NotCalled()).
			That(errFunc.MessageFormatsTo(""))
	})

	t.Run("1vs1", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.LessThan(1, 1))

		// then
		assert.
			Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got 1 <= 1"))
	})

	t.Run("1vs0", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.LessThan(1, 0))

		// then
		assert.
			Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got 1 <= 0"))
	})

}

func TestZero(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.Zero(0))

		// then
		assert.Using(t.Errorf).That(errFunc.NotCalled())
	})

	t.Run("False", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.Zero(1))

		// then
		assert.Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got 1, not zero value 0"))
	})

}

func TestNotZero(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.NotZero(1))

		// then
		assert.Using(t.Errorf).That(errFunc.NotCalled())
	})

	t.Run("False", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theval.NotZero(0))

		// then
		assert.Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got zero value 0"))
	})

}
