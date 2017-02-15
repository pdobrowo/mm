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

import "github.com/golang/go/src/pkg/sort"

type Tokens []Token

type Monomial []Token
type Polynomial []Monomial

func (lhs Monomial) Len() int {
	return len(lhs)
}

func (lhs Monomial) Swap(i, j int) {
	lhs[i], lhs[j] = lhs[j], lhs[i]
}

func (lhs Monomial) Less(i, j int) bool {
	if IsConst(lhs[i]) && IsConst(lhs[j]) {
		return Const(lhs[i]) < Const(lhs[j])
	}

	if IsVar(lhs[i]) && IsVar(lhs[j]) {
		return Var(lhs[i]) < Var(lhs[j])
	}

	if IsConst(lhs[i]) && IsVar(lhs[j]) {
		return true
	}

	if IsVar(lhs[i]) && IsConst(lhs[j]) {
		return false
	}

	panic("monomial contains a non-varconst token")
}

func NewMonomial(token Token) Monomial {
	return Monomial{token}
}

// Normalized returns a sorted monomial with at most one constant at the beginning
func (lhs Monomial) Normalized() Monomial {
	res := lhs
	sort.Sort(res)

	var mul int64
	mul = 1

	for i := 0; i < len(res); i++ {
		if !IsConst(res[i]) {
			if i > 0 {
				res = res[i-1:]
				res[0].Value = mul
			}
			break
		} else {
			mul *= Const(res[i])
		}
	}

	return res
}

func (lhs Monomial) numOfConsts() int {
	k := 0

	for i := 0; i < len(lhs); i++ {
		if IsConst(lhs[i]) {
			k++
		}
	}

	return k
}

func (lhs Monomial) IsNormalized() bool {
	return sort.IsSorted(lhs) && lhs.numOfConsts() <= 1
}

func (lhs Monomial) Pow(rhs Monomial) Monomial {
	if len(rhs) != 1 {
		panic("exponent must be a single token")
	}

	if !IsConst(rhs[0]) {
		panic("exponent must be a constant")
	}

	exponent := Const(rhs[0])

	if exponent <= 0 {
		panic("exponent must be strictly positive")
	}

	var res Monomial
	var i int64

	for _, token := range lhs {
		for i = 0; i < exponent; i++ {
			res = append(res, token)
		}
	}

	return res.Normalized()
}

func (lhs Monomial) Mul(rhs Monomial) Monomial {
	return append(lhs, rhs...).Normalized()
}

func (lhs Monomial) Add(rhs Monomial) Monomial {
	return lhs
}

func (lhs Monomial) Sub(rhs Monomial) Monomial {
	return lhs
}

func NewPolynomial(token Token) Polynomial {
	return Polynomial{NewMonomial(token)}
}

func (lhs Polynomial) Normalize() Polynomial {
	return lhs
}

func (lhs Polynomial) Pow(rhs Polynomial) Polynomial {
	return lhs
}

func (lhs Polynomial) Mul(rhs Polynomial) Polynomial {
	return lhs
}

func (lhs Polynomial) Add(rhs Polynomial) Polynomial {
	return lhs
}

func (lhs Polynomial) Sub(rhs Polynomial) Polynomial {
	return lhs
}
