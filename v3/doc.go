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

/*
Package assert provides a flexible, extensible API for making assertions.

To make a simple assertion that will panic on failure call:

	assert.UsingPanic().True(false, "Oops")

You can also format failure messages:

	assert.UsingPanic().True(0 > 1, "%d is not greater than %d", 0, 1)

UsingPanic returns an Asserter.
You can chain multiple assertions on it.

	assert.UsingPanic().
	    True(1 > 0, "%d is not greater than %d", 1, 0).
	    True(2 > 0, "%d is not greater than %d", 2, 0)

# Reusable assertions

True is good for ad hoc one-of assertions.

We provide some pre-made reusable [assertions], so you can call

	assert.UsingPanic().That(theval.Equal(got, want))

instead of

	assert.UsingPanic().True(got == want, "got %#v, not %#v", got, want)

As long as the reusable assertion is well named, the first version is easier to read.
It is also less error prone and easier to modify.

# Custom assertions

You can write your own reusable assertions as well.
Just write a function that returns a non-nil error when the assertion fails:

	func ErrIsNil(err error) error {
	    if err != nil {
			return fmt.Errorf("got unexpected non-nil error: %s", err)
		}
		retun nil
	}

You can then pass it's result to That:

	assert.UsingPanic().That(ErrIsNil(err))

You don't have to write the function in this example though.
Just use [theerr.IsNil].

# Alternative failure reactions

Depending on the situation you might want different reactions to a failed assertion.

The common case for that is to call a function that at least outputs some information.
To do that call UsingFmt.

	assert.UsingFmt(log.Panicf).That(0 > 1, "%d is not greater than %d", 0, 1)

The argument to UsingFmt has a the signature

	func(string, ...any)

Many functions and methods in the standard library have that signature.
For example:

  - testing.(*T).Errorf
  - testing.(*T).Fatalf
  - log.Printf
  - log.Panicf
  - log.Fatalf

UsingFmt will pass messages of any non-nil errors to the function.

# If for some reason you need the error itself, call Using instead.

[assertions]: https://pkg.go.dev/github.com/szabba/assert/v3/assertions
[theerr.IsNil]: https://pkg.go.dev/github.com/szabba/assert/v3/assertions/theerr#IsNil
*/
package assert
