package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func main() {
	expression, err := govaluate.NewEvaluableExpression("a + b")
	if err != nil {
		fmt.Println("Error creating expression:", err)
		return
	}

	parameters := make(map[string]interface{})
	parameters["a"] = 10
	parameters["b"] = 20

	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println("Error evaluating expression:", err)
		return
	}

	fmt.Println("Result:", result)
}