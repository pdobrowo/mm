// Copyright (c) 2017 Przemysław Dobrowolski
//
// This file is part of the math-mod, a package for symbolic manipulation
// of large algebraic expressions.
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

package math

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLexer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lexer Suite")
}

var _ = Describe("Lexer object", func() {
	var (
		infix Tokens
		err   error
	)

	Context("when whitespace is parsed", func() {
		It("should succeed", func() {
			infix, err = ParseInfixString("  \t  x \n +   \v 1\f     \r")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("when different tokens are parsed", func() {
		It("should succeed", func() {
			infix, err = ParseInfixString("x+y* z ^ w + ( x (y + z)^3 + 7 )")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("when integer is on its own", func() {
		It("should succeed", func() {
			infix, err = ParseInfixString("7")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(infix).To(Equal(Tokens{NewConst(7)}))
		})
	})

	Context("when variable is on its own", func() {
		It("should succeed", func() {
			infix, err = ParseInfixString("x")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(infix).To(Equal(Tokens{NewVar("x")}))
		})
	})

	Context("when operator is on its own", func() {
		It("should succeed", func() {
			infix, err = ParseInfixString("+")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(infix).To(Equal(Tokens{NewPlus()}))
		})
	})

	Context("when bracket is on its own", func() {
		It("should succeed", func() {
			infix, err = ParseInfixString("(")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(infix).To(Equal(Tokens{NewOpen()}))
		})
	})
})
