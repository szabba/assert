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
Package assert provides a core, minimal API for making assertions.

To make a simple assertion that will panic on failure call:

	assert.UsingPanic().That(false, "Oops")

You can also format failure messages:

	assert.UsingPanic().That(0 > 1, "%d is not greater than %d", 0, 1)

UsingPanic returns an Asserter.
You can chain multiple assertions on it.

	assert.UsingPanic().
	    That(1 > 0, "%d is not greater than %d", 1, 0).
	    That(2 > 0, "%d is not greater than %d", 2, 0)

# Alternative failure reactions

Depending on the situation you might want different reactions to a failed assertion.
To do that call Using with an ErrorFunc.

	assert.Using(log.Panicf).That(0 > 1, "%d is not greater than %d", 0, 1)

An ErrorFunc specifies what the reaction to failure should be.
When Using is passed a nil ErrorFunc, it behaves the same as UsingPanic.

You can use many functions and methods in the standard library as ErrorFuncs.
For example:

  - testing.(*T).Errorf
  - testing.(*T).Fatalf
  - log.Printf
  - log.Panicf
  - log.Fatalf

# Reusable assertions

We provide some pre-made reusable [assertions], so you can call

	assert.UsingPanic().That(equal.Values(got, want))

instead of

	assert.UsingPanic().That(got == want, "got %#v, not %#v", got, want)

As long as the reusable assertion is well named, the first version is easier to read.
The first version is also less error prone and easier to modify.
In more complex cases a reusable assertion can also provide more detailed error messages.

Reusable assertions rely on a feature of Go that is used relatively rarely.
You can read about it in the [Calls] section of the Go language specification.

# Custom assertions

You can write your own reusable assertions as well.
Just write a function that returns the arguments you would pass to That:

	func ErrIsNil(err error) (bool, string) {
	    return err == nil, fmt.Sprintf("unexpected error: %s", err)
	}

You can then pass it's result to That instead of writing the arguments out:

	assert.UsingPanic().That(ErrIsNil(err))

We could have written ErrIsNil this way as well:

	func ErrIsNil(err error) (bool, string, any) {
	    return err == nil, "unexpected error: %s, any
	}

We recommend that custom assertions return (bool, string).
All the reusable assertions we provide are written this way.
This

  - makes them easier to recognize,
  - makes them easier to modify, and
  - allows us to provide better error messages for complex cases.

You don't have to write the function in this example though.
Just use [theerr.IsNil].

[assertions]: https://pkg.go.dev/github.com/szabba/assert/v2/assertions
[Calls]: https://go.dev/ref/spec#Calls
[theerr.IsNil]: https://pkg.go.dev/github.com/szabba/assert/v2/assertions/theerr#IsNil
*/
package assert
