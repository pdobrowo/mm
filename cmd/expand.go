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

package cmd

import (
	"github.com/pdobrowo/mm/math"
	"github.com/spf13/cobra"
)

func expandCmdRun(cmd *cobra.Command, args []string) error {
	tokens, err := parseArgs(args, true)

	if err != nil {
		return err
	}

	var stack []math.Polynomial

	for _, token := range tokens {
		if math.IsVarConst(token) {
			stack = append(stack, math.NewPolynomial(token))
			continue
		}

		if !math.IsOp(token) {
			panic("token is not an operator")
		}

		if len(stack) < 2 {
			panic("stack is tool small")
		}

		lhs := stack[len(stack)-1]
		rhs := stack[len(stack)-1]

		stack = stack[:len(stack)-2]

		switch token.Kind {
		case math.KindOpPow:
			stack = append(stack, lhs.Pow(rhs))
		case math.KindOpMul:
			stack = append(stack, lhs.Mul(rhs))
		case math.KindOpAdd:
			stack = append(stack, lhs.Add(rhs))
		case math.KindOpSub:
			stack = append(stack, lhs.Sub(rhs))
		}
	}

	if len(stack) != 1 {
		panic("stack should contain one token")
	}

	// print

	return nil
}

// expandCmd represents the expand command
var expandCmd = &cobra.Command{
	Use:   "expand",
	Short: "Expand an algebraic expression",
	Long:  `Expands a given expression removing all brackets`,
	RunE:  expandCmdRun,
}

func init() {
	RootCmd.AddCommand(expandCmd)
}
