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

package theslice_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/szabba/assert/v2"
	"github.com/szabba/assert/v2/assertions/assertiontesting"

	"github.com/szabba/assert/v2/assertions/theslice"
)

func TestEmpty(t *testing.T) {

	t.Run("True/Nil", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theslice.Empty[[]int](nil))

		// then
		assert.Using(t.Errorf).That(errFunc.NotCalled())
	})

	t.Run("True/NotNil", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc
		s := []int{}

		// when
		assert.Using(errFunc.Record).That(theslice.Empty(s))

		// then
		assert.Using(t.Errorf).That(errFunc.NotCalled())
	})

	t.Run("False", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc
		s := []int{1, 2}

		// when
		assert.Using(errFunc.Record).That(theslice.Empty(s))

		// then
		assert.Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got non-empty slice []int{1, 2}"))
	})

}

func TestNotEmpty(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc
		s := []int{1, 2}

		// when
		assert.Using(errFunc.Record).That(theslice.NotEmpty(s))

		// then
		assert.Using(t.Errorf).That(errFunc.NotCalled())
	})

	t.Run("False/NotNil", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc
		s := []int{}

		// when
		assert.Using(errFunc.Record).That(theslice.NotEmpty(s))

		// then
		assert.Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got empty slice []int{}"))
	})

	t.Run("False/Nil", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(theslice.NotEmpty[[]int](nil))

		// then
		assert.Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got empty slice []int(nil)"))
	})

}

func TestEqual(t *testing.T) {

	nameGiven := func(got, want []int) string {
		mightHaveSpaces := fmt.Sprintf("Got%#vWant%#v", got, want)
		return strings.ReplaceAll(mightHaveSpaces, " ", "")
	}

	okCases := [][]int{
		nil,
		{},
		{0},
		{0, 1},
	}

	for _, tt := range okCases {
		name := nameGiven(tt, tt)

		t.Run(name, func(t *testing.T) {
			// given
			var onErr assertiontesting.ErrFunc

			// when
			assert.Using(onErr.Record).That(theslice.Equal(tt, tt))

			// then
			assert.Using(t.Errorf).That(onErr.NotCalled())
		})
	}

	oopsCases := []struct {
		Got, Want []int
		Message   string
	}{
		{
			Got:     nil,
			Want:    []int{},
			Message: "got nil, not []int{}",
		},
		{
			Got:     []int{},
			Want:    nil,
			Message: "got []int{}, not nil",
		},
		{
			Got:     []int{},
			Want:    []int{0},
			Message: "got []int{} (of length 0), not []int{0} (of length 1)",
		},
		{
			Got:     []int{1},
			Want:    []int{2},
			Message: "got slice []int{1}, not []int{2}: element at position 0 is 1, not 2",
		},
		{
			Got:     []int{1, 2},
			Want:    []int{1, 3},
			Message: "got slice []int{1, 2}, not []int{1, 3}: element at position 1 is 2, not 3",
		},
		{
			Got:     []int{1, 2},
			Want:    []int{3, 4},
			Message: "got slice []int{1, 2}, not []int{3, 4}: element at position 0 is 1, not 3; element at position 1 is 2, not 4",
		},
	}

	for _, tt := range oopsCases {

		name := nameGiven(tt.Got, tt.Want)

		t.Run(name, func(t *testing.T) {
			// given
			var onErr assertiontesting.ErrFunc

			// when
			assert.
				Using(onErr.Record).
				That(theslice.Equal(tt.Got, tt.Want))

			// then
			assert.
				Using(t.Errorf).
				That(onErr.Called()).
				That(onErr.MessageFormatsTo(tt.Message))
		})
	}
}

func TestNotEqual(t *testing.T) {

	nameGiven := func(got, want []int) string {
		mightHaveSpaces := fmt.Sprintf("Got%#vWant%#v", got, want)
		return strings.ReplaceAll(mightHaveSpaces, " ", "")
	}

	oopsCases := []struct {
		Slice   []int
		Message string
	}{
		{
			Slice:   nil,
			Message: "got unwanted value []int(nil)",
		},
		{
			Slice:   []int{},
			Message: "got unwanted value []int{}",
		},
		{
			Slice:   []int{0},
			Message: "got unwanted value []int{0}",
		},
		{
			Slice:   []int{1, 2},
			Message: "got unwanted value []int{1, 2}",
		},
	}

	for _, tt := range oopsCases {
		name := nameGiven(tt.Slice, tt.Slice)

		t.Run(name, func(t *testing.T) {
			// given
			var errFunc assertiontesting.ErrFunc

			// when
			assert.Using(errFunc.Record).
				That(theslice.NotEqual(tt.Slice, tt.Slice))

			// then
			assert.Using(t.Errorf).
				That(errFunc.Called()).
				That(errFunc.MessageFormatsTo(tt.Message))
		})
	}

	okCases := []struct {
		Got, Want []int
	}{
		{Got: nil, Want: []int{}},
		{Got: []int{}, Want: nil},
		{Got: []int{1}, Want: []int{2, 3}},
		{Got: []int{2, 3}, Want: []int{1}},
		{Got: []int{0}, Want: []int{1}},
		{Got: []int{1}, Want: []int{0}},
	}

	for _, tt := range okCases {
		name := nameGiven(tt.Got, tt.Want)

		t.Run(name, func(t *testing.T) {
			// given
			var errFunc assertiontesting.ErrFunc

			// when
			assert.Using(errFunc.Record).
				That(theslice.NotEqual(tt.Got, tt.Want))

			// then
			assert.Using(t.Errorf).
				That(errFunc.NotCalled())
		})
	}

}

func TestLength(t *testing.T) {

	okCases := map[string]struct {
		Slice  []int
		Length int
	}{
		"True/Nil":   {Slice: nil, Length: 0},
		"True/Empty": {Slice: []int{}, Length: 0},
		"True/1":     {Slice: []int{1}, Length: 1},
		"True/2":     {Slice: []int{1, 2}, Length: 2},
		"True/3":     {Slice: []int{1, 2, 3}, Length: 3},
	}

	for name, tt := range okCases {
		t.Run(name, func(t *testing.T) {
			// given
			var errFunc assertiontesting.ErrFunc

			// when
			assert.Using(errFunc.Record).
				That(theslice.Length(tt.Slice, tt.Length))

			// then
			assert.Using(t.Errorf).That(errFunc.NotCalled())
		})
	}

	oopsCases := map[string]struct {
		Slice   []int
		Length  int
		Message string
	}{
		"False/NilHasLenght1": {
			Slice:   nil,
			Length:  1,
			Message: "got slice of length 0, not 1",
		},
		"False/EmptyHasLength1": {
			Slice:   []int{},
			Length:  1,
			Message: "got slice of length 0, not 1",
		},
		"False/Make1HasLength2": {
			Slice:   make([]int, 1),
			Length:  2,
			Message: "got slice of length 1, not 2",
		},
		"False/Make2HasLength3": {
			Slice:   make([]int, 2),
			Length:  3,
			Message: "got slice of length 2, not 3",
		},
	}

	for name, tt := range oopsCases {
		t.Run(name, func(t *testing.T) {
			// given
			var errFunc assertiontesting.ErrFunc

			// when
			assert.Using(errFunc.Record).
				That(theslice.Length(tt.Slice, tt.Length))

			// then
			assert.Using(t.Errorf).
				That(errFunc.Called()).
				That(errFunc.MessageFormatsTo(tt.Message))
		})
	}

}

func TestLengthNot(t *testing.T) {

	okCases := map[string]struct {
		Slice  []int
		Length int
	}{
		"True/NilNotLength1": {
			Slice:  nil,
			Length: 1,
		},
		"True/EmptyNotLength1": {
			Slice:  []int{},
			Length: 1,
		},
		"True/Make1NotLength2": {
			Slice:  make([]int, 1),
			Length: 2,
		},
		"True/Make2NotLength3": {
			Slice:  make([]int, 2),
			Length: 3,
		},
	}

	for name, tt := range okCases {
		t.Run(name, func(t *testing.T) {
			// given
			var errFunc assertiontesting.ErrFunc

			// when
			assert.Using(errFunc.Record).
				That(theslice.LengthNot(tt.Slice, tt.Length))

			// then
			assert.Using(t.Errorf).That(errFunc.NotCalled())
		})
	}

	oopsCases := map[string]struct {
		Slice   []int
		Length  int
		Message string
	}{
		"False/NilNotLength0": {
			Slice:   nil,
			Length:  0,
			Message: "got slice of length 0",
		},
		"False/EmptyNotLength0": {
			Slice:   []int{},
			Length:  0,
			Message: "got slice of length 0",
		},
		"False/Make1NotLength1": {
			Slice:   make([]int, 1),
			Length:  1,
			Message: "got slice of length 1",
		},
		"False/Make2NotLength2": {
			Slice:   make([]int, 2),
			Length:  2,
			Message: "got slice of length 2",
		},
	}

	for name, tt := range oopsCases {
		t.Run(name, func(t *testing.T) {
			// given
			var errFunc assertiontesting.ErrFunc

			// when
			assert.Using(errFunc.Record).
				That(theslice.LengthNot(tt.Slice, tt.Length))

			// then
			assert.Using(t.Errorf).
				That(errFunc.Called()).
				That(errFunc.MessageFormatsTo(tt.Message))
		})
	}

}
