package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"strconv"
	"errors"
)


func main() {
	
	//Context
	resultCh := make(chan string, 10)
	depthCh := make(chan string, 10)
	errCh := make(chan error)
	doneCh := make(chan struct{})

	// Goroutine A
	go func() {
		exprA := "234+34/34"
		valueA, err := evaluateExpression(exprA)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- valueA
		depthCh <- valueA
	}()

	// Goroutine B
	go func() {
		valueA := <-depthCh // Wait for A to provide value
		exprB := "a+234/34.3"
		valueB, err := evaluateExpressionWithVariable(exprB, "a", valueA)
		if err != nil {
			fmt.Println("Error evaluating expression in B:", err)
			return
		}

		resultCh <- valueB
		
	}()

	// Goroutine C
	go func() {
		valueB := <-depthCh // Wait for B to provide value
		exprC := "a+b"
		valueC, err := evaluateExpressionWithVariable(exprC, "b", valueB)
		if err != nil {
			fmt.Println("Error evaluating expression in C:", err)
			return
		}
		resultCh <- valueC
		doneCh <- struct{}{}
	}

	for results := range resultCh {
		fmt.Println("Result:", results)
	}


	
}

func evaluateExpression(expr string) (string, error) {
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return "", err
	}

	parameters := make(map[string]interface{})
	result, err := expression.Evaluate(parameters)
	if err != nil {
		return "", err
	}

	switch v := result.(type) {
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case string:
		// If the result is already a string, return it directly
		return v, nil
	default:
		return "", errors.New("unsupported result type")
	}
}

func evaluateExpressionWithVariable(expr, variableName, variableValue string) (string, error) {
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return "", err
	}

	parameters := make(map[string]interface{})
	parameters[variableName] = variableValue
	result, err := expression.Evaluate(parameters)
	if err != nil {
		return "", err
	}

	switch v := result.(type) {
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case string:
		// If the result is already a string, return it directly
		return v, nil
	default:
		return "", errors.New("unsupported result type")
	}
}