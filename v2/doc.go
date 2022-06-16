// MIT License
//
// Copyright (c) 2018 Karol Marcjan
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

// Package assert provides a core, minimal API for making assertions.
//
// To make a simple assertion that will panic on failure call:
//
//     assert.UsingPanic().That(false, "Oops")
//
// You can also format failure messages:
//
//     assert.UsingPanic().That(0 > 1, "%d is not greater than %d", 0, 1)
//
// UsingPanic returns an Asserter.
// You can chain multiple assertions on it.
//
//     assert.UsingPanic().
//         That(1 > 0, "%d is not greater than %d", 1, 0).
//         That(2 > 0, "%d is not greater than %d", 2, 0)
//
// Alternative failure reactions
//
// Depending on the situation you might want different reactions to a failed assertion.
// To do that call Using with an ErrorFunc.
//
//     assert.Using(log.Panicf).That(0 > 1, "%d is not greater than %d", 0, 1)
//
// An ErrorFunc specifies what the reaction to failure should be.
// When Using is passed a nil ErrorFunc, it behaves the same as UsingPanic.
//
// You can use many functions and methods in the standard library as ErrorFuncs.
// For example:
//
//     - testing.(*T).Errorf
//     - testing.(*T).Fatalf
//     - log.Printf
//     - log.Panicf
//     - log.Fatalf
//
// Reusable assertions
//
// For a one time assertion it is OK to write out the arguments to That.
// If you repeat them many times - with little or no variation - it quickly becomes error prone.
//
// Instead it is better to write a function that returns the arguments you would pass to That:
//
//     func NilError(err error) (bool, string, any) {
//         return err == nil, "unexpected error: %s", err
//     }
//
// You can then pass it's result to That instead of writing the arguments out:
//
//     assert.UsingPanic().That(NilError(err))
//
// This relies on a feature of Go that is used relatively rarely.
// You can read about it in the [Calls] section of the Go language specification.
//
// [Calls]: https://go.dev/ref/spec#Calls
package assert
