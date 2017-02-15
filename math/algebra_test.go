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

func TestAlgebra(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Algebra Suite")
}

var _ = Describe("Algebra object", func() {
	Context("when monomials are normalized", func() {
		It("should succeed", func() {
			m := Monomial{
				NewVar("x"),
				NewVar("z"),
				NewVar("a"),
				NewConst(7),
				NewConst(2),
				NewVar("m"),
			}

			m = m.Normalized()

			Expect(m.IsNormalized()).To(Equal(true))

			Expect(m).To(Equal(Monomial{
				NewConst(14),
				NewVar("a"),
				NewVar("m"),
				NewVar("x"),
				NewVar("z"),
			}))
		})
	})

	Context("when monomials are powered", func() {
		It("should succeed", func() {
			lhs := Monomial{
				NewVar("x"),
				NewVar("y"),
				NewVar("z"),
			}

			rhs := Monomial{
				NewConst(3),
			}

			Expect(lhs.Pow(rhs)).To(Equal(Monomial{
				NewVar("x"),
				NewVar("x"),
				NewVar("x"),
				NewVar("y"),
				NewVar("y"),
				NewVar("y"),
				NewVar("z"),
				NewVar("z"),
				NewVar("z"),
			}))
		})
	})

	Context("when monomials are multiplies", func() {
		It("should succeed", func() {
			lhs := Monomial{
				NewConst(2),
				NewVar("x"),
				NewVar("y"),
				NewVar("z"),
			}

			rhs := Monomial{
				NewConst(3),
				NewVar("a"),
			}

			Expect(lhs.Mul(rhs)).To(Equal(Monomial{
				NewConst(6),
				NewVar("a"),
				NewVar("x"),
				NewVar("y"),
				NewVar("z"),
			}))
		})
	})
})
