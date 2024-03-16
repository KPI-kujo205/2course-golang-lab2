package main

import (
	"bufio"
	"flag"
	"fmt"
	lab2 "github.com/KPI-kujo205/2course-golang-lab2"
	"io"
	"os"
	"strings"
)

var (
	inputExpression     = flag.String("e", "", "Expression to compute")
	inputExpressionFile = flag.String("f", "", "Path to input expression file")
	outputFile          = flag.String("o", "", "Path to output file")
)

func main() {
	flag.Parse()

	if err := validateFlags(); err != nil {
		exitWithError(err)
	}

	input, err := getInputReader()
	if err != nil {
		exitWithError(err)
	}
	defer func() {
		if closer, ok := input.(io.Closer); ok {
			if err := closer.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "error closing input: %v\n", err)
			}
		}
	}()

	output, err := getOutputWriter()
	if err != nil {
		exitWithError(err)
	}
	defer func() {
		if closer, ok := output.(io.Closer); ok {
			if err := closer.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "error closing output: %v\n", err)
			}
		}
	}()

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	if err := handler.Compute(); err != nil {
		exitWithError(err)
	}

	if writer, ok := output.(*bufio.Writer); ok {
		if err := writer.Flush(); err != nil {
			fmt.Fprintln(os.Stderr, "error flushing output writer:", err)
		}
	}
}

func validateFlags() error {
	if *inputExpression == "" && *inputExpressionFile == "" {
		return fmt.Errorf("input expression or input file path must be provided")
	} else if *inputExpression != "" && *inputExpressionFile != "" {
		return fmt.Errorf("only one flag available: -e or -f")
	}
	return nil
}

func getInputReader() (io.Reader, error) {
	if *inputExpression != "" {
		return strings.NewReader(*inputExpression), nil
	} else if *inputExpressionFile != "" {
		file, err := os.Open(*inputExpressionFile)
		if err != nil {
			return nil, fmt.Errorf("error opening input file: %v", err)
		}
		return bufio.NewReader(file), nil
	}
	return nil, nil
}

func getOutputWriter() (io.Writer, error) {
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			return nil, fmt.Errorf("error creating output file: %v", err)
		}
		return bufio.NewWriter(file), nil
	}
	return os.Stdout, nil
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
