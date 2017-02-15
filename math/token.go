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
	KindConst = iota

	// variable
	KindVar

	// operators
	KindOpAdd
	KindOpSub
	KindOpMul
	KindOpPow

	// parenthesis
	KindParOpen
	KindParClose
)

type Kind int

var OperProps = map[Kind]struct {
	prec       int  // precedence
	rightAssoc bool // right-associativity
}{
	KindOpPow: {4, true},
	KindOpMul: {3, false},
	KindOpAdd: {2, false},
	KindOpSub: {2, false},
}

type Token struct {
	Kind  Kind
	Value interface{}
}

func IsOp(token Token) bool {
	switch token.Kind {
	case KindOpPow, KindOpMul, KindOpAdd, KindOpSub:
		return true
	default:
		return false
	}
}

func IsPar(token Token) bool {
	switch token.Kind {
	case KindParOpen, KindParClose:
		return true
	default:
		return false
	}
}

func IsConst(token Token) bool {
	return token.Kind == KindConst
}

func IsVar(token Token) bool {
	return token.Kind == KindVar
}

func IsVarConst(token Token) bool {
	return IsVar(token) || IsConst(token)
}

func Var(token Token) string {
	return token.Value.(string)
}

func Const(token Token) int64 {
	return token.Value.(int64)
}

func NewConst(value int64) Token {
	return Token{
		Kind:  KindConst,
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
	return Token{Kind: KindOpAdd}
}

func NewMinus() Token {
	return Token{Kind: KindOpSub}
}

func NewMul() Token {
	return Token{Kind: KindOpMul}
}

func NewPow() Token {
	return Token{Kind: KindOpPow}
}

func NewOpen() Token {
	return Token{Kind: KindParOpen}
}

func NewClose() Token {
	return Token{Kind: KindParClose}
}

func (token Token) String() string {
	switch token.Kind {
	case KindConst:
		return fmt.Sprintf("%d", Const(token))
	case KindVar:
		return Var(token)
	case KindOpAdd:
		return "+"
	case KindOpSub:
		return "-"
	case KindOpMul:
		return "*"
	case KindOpPow:
		return "^"
	case KindParOpen:
		return "("
	case KindParClose:
		return ")"
	}

	panic("invalid token kind")
}
