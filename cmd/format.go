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
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/pdobrowo/mm/math"
	"github.com/spf13/cobra"
)

var postfixFlag *bool

func checkFlags() error {
	return nil
}

func formatCmdRun(cmd *cobra.Command, args []string) error {
	var reader io.Reader
	var err error

	if err = checkFlags(); err != nil {
		return err
	}

	switch len(args) {
	case 0:
		reader = bufio.NewReader(os.Stdin)
	case 1:
		reader, err = os.Open(args[0])

		if err != nil {
			return fmt.Errorf("failed to open file: %v", args[0])
		}
	default:
		return fmt.Errorf("invalid number of arguments: %d", len(args))
	}

	infix, err := math.ParseInfix(reader)

	if err != nil {
		return err
	}

	infix = math.ImplicitOperMul(infix)

	if *postfixFlag == true {
		postfix := math.ToPostfix(infix)

		fmt.Println(postfix)
	} else {
		fmt.Println(infix)
	}

	return nil
}

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format an algebraic expression",
	Long: `Formatting a large algebraic expression
is usually a first step to discover its properties
and possible simplifications.`,
	RunE: formatCmdRun,
}

func init() {
	RootCmd.AddCommand(formatCmd)

	postfixFlag = formatCmd.PersistentFlags().Bool("postfix", false, "Use postfix (RPN) format")
}
