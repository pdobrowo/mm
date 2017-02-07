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
	"bufio"
	"io"

	"github.com/golang/go/src/pkg/strconv"
)

const (
	KindInt = iota

	KindVar

	KindPlus
	KindMinus
	KindMul
	KindPow

	KindOpen
	KindClose
)

type Token struct {
	kind  int
	value interface{}
}

func (token Token) String() string {
	switch token.kind {
	case KindInt:
		return string(token.value.(int))
	case KindVar:
		return token.value.(string)
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

func Parse(reader io.Reader) ([]Token, error) {
	tokens := []Token{}
	scanner := bufio.NewScanner(reader)

	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		i := 0

		// omit whitespace
	omit_whitespace:
		for i < len(data) {
			switch data[i] {
			case '\t', '\n', '\v', '\f', '\r', ' ':
				i++
				continue
			default:
				break omit_whitespace
			}
		}

		if i == len(data) {
			return i, nil, nil
		}

		// scan one-character tokens
		switch data[i] {
		case '+', '-', '*', '^', '(', ')':
			return i + 1, data[i : i+1], nil
		}

		// scan integer or variable
		j := i

		for j < len(data) {
			if (data[j] >= 'a' && data[j] <= 'z') || (data[j] >= 'A' && data[j] <= 'Z') || (data[j] >= '0' && data[j] <= '9') {
				j++
			} else {
				break
			}
		}

		if j < len(data) {
			return j, data[i:j], nil
		}

		return i, nil, nil
	})

next_token:
	for scanner.Scan() {
		raw := scanner.Text()
		switch raw {
		case "+":
			tokens = append(tokens, Token{
				kind:  KindPlus,
				value: nil,
			})

			continue

		case "-":
			tokens = append(tokens, Token{
				kind:  KindMinus,
				value: nil,
			})

			continue

		case "*":
			tokens = append(tokens, Token{
				kind:  KindMul,
				value: nil,
			})

			continue

		case "^":
			tokens = append(tokens, Token{
				kind:  KindPow,
				value: nil,
			})

			continue

		case "(":
			tokens = append(tokens, Token{
				kind:  KindOpen,
				value: nil,
			})

			continue

		case ")":
			tokens = append(tokens, Token{
				kind:  KindClose,
				value: nil,
			})

			continue
		}

		// integer or variable
		for r := range raw {
			if r < '0' || r > '9' {
				tokens = append(tokens, Token{
					kind:  KindVar,
					value: raw,
				})

				continue next_token
			}
		}

		i, err := strconv.ParseInt(raw, 10, 0)

		if err != nil {
			return nil, err
		}

		tokens = append(tokens, Token{
			kind:  KindInt,
			value: i,
		})
	}

	return tokens, nil
}
