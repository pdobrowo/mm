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
	"fmt"
)

const (
	// numbers
	KindInt = iota

	// variable
	KindVar

	// operators
	KindPlus
	KindMinus
	KindMul
	KindPow

	// brackets
	KindOpen
	KindClose
)

type Kind int

var OperProps = map[Kind]struct {
	prec       int  // precedence
	rightAssoc bool // right-associativity
}{
	KindPow:   {4, true},
	KindMul:   {3, false},
	KindPlus:  {2, false},
	KindMinus: {2, false},
}

type Token struct {
	Kind  Kind
	Value interface{}
}

func NewInt(value int64) Token {
	return Token{
		Kind:  KindInt,
		Value: value,
	}
}

func NewVar(value string) Token {
	return Token{
		Kind:  KindVar,
		Value: value,
	}
}

func NewPlus() Token {
	return Token{Kind: KindPlus}
}

func NewMinus() Token {
	return Token{Kind: KindMinus}
}

func NewMul() Token {
	return Token{Kind: KindMul}
}

func NewPow() Token {
	return Token{Kind: KindPow}
}

func NewOpen() Token {
	return Token{Kind: KindOpen}
}

func NewClose() Token {
	return Token{Kind: KindClose}
}

func (token Token) String() string {
	switch token.Kind {
	case KindInt:
		return fmt.Sprintf("%d", token.Value.(int64))
	case KindVar:
		return token.Value.(string)
	case KindPlus:
		return "+"
	case KindMinus:
		return "-"
	case KindMul:
		return "*"
	case KindPow:
		return "^"
	case KindOpen:
		return "("
	case KindClose:
		return ")"
	}

	panic("invalid token kind")
}
