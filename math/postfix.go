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

func ToPostfix(infix Tokens) (postfix Tokens) {
	var stack Tokens

	for _, token := range infix {
		switch token.Kind {
		case KindParOpen:
			stack = append(stack, token)
		case KindParClose:
			var op Token
			for {
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op.Kind == KindParOpen {
					break
				}
				postfix = append(postfix, op)
			}
		default:
			if operPropA, isOperA := OperProps[token.Kind]; isOperA {
				for len(stack) > 0 {
					oper := stack[len(stack)-1]
					if operPropB, isOperB := OperProps[oper.Kind]; !isOperB || operPropA.prec > operPropB.prec || operPropA.prec == operPropB.prec && operPropA.rightAssoc {
						break
					}
					stack = stack[:len(stack)-1]
					postfix = append(postfix, oper)
				}
				stack = append(stack, token)
			} else {
				postfix = append(postfix, token)
			}
		}
	}
	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return
}
