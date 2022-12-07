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

package theerr_test

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/szabba/assert/v2"
	"github.com/szabba/assert/v2/assertions/assertiontesting"
	theerr "github.com/szabba/assert/v2/assertions/theerr"
)

func TestIsNil(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theerr.IsNil(nil))

		// then
		assert.Using(t.Errorf).That(errFunc.NotCalled())
	})

	t.Run("False", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theerr.IsNil(io.EOF))

		// then
		assert.
			Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("unexpected error: EOF"))
	})

}

func TestIs(t *testing.T) {

	t.Run("True/Nil", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theerr.Is(nil, nil))

		// then
		assert.Using(t.Errorf).That(errFunc.NotCalled())
	})

	t.Run("True/Exact", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theerr.Is(io.EOF, io.EOF))

		// then
		assert.Using(t.Errorf).That(errFunc.NotCalled())
	})

	t.Run("True/Wrapped", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		wrappedErr := fmt.Errorf("failed: %w", io.EOF)

		// when
		assert.Using(errFunc.Record).That(theerr.Is(wrappedErr, io.EOF))

		// then
		assert.Using(t.Errorf).That(errFunc.NotCalled())
	})

	t.Run("False", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		newErr := errors.New("oops")

		// when
		assert.Using(errFunc.Record).That(theerr.Is(newErr, io.EOF))

		// then
		assert.Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo(`got "oops", not "EOF"`))
	})

}

func TestIsA(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc
		err := new(os.PathError)

		// when
		assert.Using(errFunc.Record).
			That(theerr.IsA[*os.PathError](err))

		// then
		assert.Using(t.Errorf).That(errFunc.NotCalled())
	})

	t.Run("False", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc
		err := new(os.SyscallError)

		// when
		assert.Using(errFunc.Record).
			That(theerr.IsA[*os.PathError](err))

		// then
		assert.Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got error of type *os.SyscallError, not *fs.PathError"))
	})

}
