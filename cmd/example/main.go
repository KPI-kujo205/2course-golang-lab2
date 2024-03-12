package main

import (
	"flag"
	"fmt"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression    = flag.String("e", "", "Expression to compute")
	inputExpresionFile = flag.String("f", "", "Path to input expression file")
	outputFile         = flag.String("o", "", "Path to output file")
)

func main() {
	flag.Parse()

	if *inputExpression != "" && *inputExpresionFile != "" {
		panic("Only one flag available: -e or -f")
	}

	fmt.Println(*inputExpression, *inputExpresionFile, *outputFile) // test
	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, _ := lab2.PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
