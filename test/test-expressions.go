package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func main() {
	expression, err := govaluate.NewEvaluableExpression("a + b - c/1.8")
	if err != nil {
		fmt.Println("Error creating expression:", err)
		return
	}

	parameters := make(map[string]interface{})
	parameters["a"] = 10
	parameters["b"] = 20
	parameters["c"] = 6

	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println("Error evaluating expression:", err)
		return
	}

	fmt.Println("Result:", result)
}