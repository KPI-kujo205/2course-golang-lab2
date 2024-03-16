package lab2

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ComputeHandler is constructed with input io.Reader and output io.Writer.
// Its Compute() method reads the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(ch.Input)
	if err != nil {
		return err
	}

	lines := strings.Split(buffer.String(), "\n")
	for _, line := range lines {
		inputString := strings.TrimSpace(line)

		postfixExpr, err := PrefixToPostfix(inputString)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(ch.Output, postfixExpr)
		if err != nil {
			return err
		}
	}

	return nil
}
