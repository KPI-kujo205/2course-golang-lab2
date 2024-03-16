package lab2

import (
	"bytes"
	"gopkg.in/check.v1"
	"strings"
)

type ComputeHandlerSuite struct{}

var _ = check.Suite(&ComputeHandlerSuite{})

func (s *ComputeHandlerSuite) TestValidExpression(c *check.C) {
	validExpression := "+ + 1 * 2 3 9"
	expectedExpression := "1 2 3 * + 9 +\n"
	outputBuffer := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader(validExpression),
		Output: outputBuffer,
	}

	err := handler.Compute()
	c.Assert(err, check.IsNil)
	c.Assert(outputBuffer.String(), check.Equals, expectedExpression)
}

func (s *ComputeHandlerSuite) TestInvalidCharactersInput(c *check.C) {
	invalidCharactersInput := "wrong characters"
	outputBuffer := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader(invalidCharactersInput),
		Output: outputBuffer,
	}

	err := handler.Compute()
	c.Assert(err, check.NotNil)
}

func (s *ComputeHandlerSuite) TestInvalidExpression(c *check.C) {
	outputBuffer := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader("+ + 1"),
		Output: outputBuffer,
	}
	err := handler.Compute()

	c.Assert(err, check.NotNil)
}
