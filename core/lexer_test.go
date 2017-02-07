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

package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"strings"
)

func TestLexer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lexer Suite")
}

func parseString(eqn string) ([]Token, error) {
	return Parse(strings.NewReader(eqn))
}

var _ = Describe("Lexer Object", func() {
	var (
		tokens []Token
		err error
	)

	BeforeEach(func() {

	})

	Context("when whitespace is parsed", func() {
		It("should succeed", func() {
			tokens, err = parseString("  \t  x \n +   \v 1\f     \r")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("when different tokens are parsed", func() {
		It("should succeed", func() {
			tokens, err = parseString("x+y* z ^ w + ( x (y + z)^3 + 7 )")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
