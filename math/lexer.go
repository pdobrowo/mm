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
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ParseInfix(reader io.Reader) (Tokens, error) {
	tokens := Tokens{}
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
			return i + 1, data[i: i+1], nil
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

		return j, data[i:j], nil
	})

	for scanner.Scan() {
		raw := scanner.Text()
		switch raw {
		case "+":
			tokens = append(tokens, NewPlus())
			continue

		case "-":
			tokens = append(tokens, NewMinus())
			continue

		case "*":
			tokens = append(tokens, NewMul())
			continue

		case "^":
			tokens = append(tokens, NewPow())
			continue

		case "(":
			tokens = append(tokens, NewOpen())
			continue

		case ")":
			tokens = append(tokens, NewClose())
			continue
		}

		// integer or variable
		i, err := strconv.ParseInt(raw, 10, 64)

		if err != nil {
			tokens = append(tokens, NewVar(raw))
		} else {
			tokens = append(tokens, NewConst(i))
		}
	}

	return tokens, nil
}

func ParseInfixString(infix string) (Tokens, error) {
	return ParseInfix(strings.NewReader(infix))
}
