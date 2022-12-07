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

package equal_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/szabba/assert/v2"
	"github.com/szabba/assert/v2/assertions/assertiontesting"

	"github.com/szabba/assert/v2/assertions/equal"
)

func TestValues(t *testing.T) {

	t.Run("True", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(equal.Values(0, 0))

		// then
		assert.
			Using(t.Errorf).
			That(errFunc.NotCalled())
	})

	t.Run("False", func(t *testing.T) {
		// given
		var errFunc assertiontesting.ErrFunc

		// when
		assert.Using(errFunc.Record).That(equal.Values(0, 1))

		// then
		assert.
			Using(t.Errorf).
			That(errFunc.Called()).
			That(errFunc.MessageFormatsTo("got 0, not 1"))
	})

}

func TestSlices(t *testing.T) {

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

	for _, slice := range okCases {
		name := nameGiven(slice, slice)

		t.Run(name, func(t *testing.T) {
			// given
			var onErr assertiontesting.ErrFunc

			// when
			assert.
				Using(onErr.Record).
				That(equal.Slices(slice, slice))

			// then
			assert.
				Using(t.Errorf).
				That(onErr.NotCalled())
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
				That(equal.Slices(tt.Got, tt.Want))

			// then
			assert.
				Using(t.Errorf).
				That(onErr.Called()).
				That(onErr.MessageFormatsTo(tt.Message))
		})
	}
}
