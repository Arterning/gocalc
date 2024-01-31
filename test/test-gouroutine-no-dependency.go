package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"strconv"
	"sync"
)

func main() {
	expressions := []string{"a + b", "c * d", "e - f", "(a + b) * c"} // Example expressions
	variables := map[string]int{"a": 10, "b": 20, "c": 5, "d": 8, "e": 100, "f": 60} // Example variable values

	j := 5 // Maximum number of goroutines

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, j)

	results := make(map[string]string)
	var mu sync.Mutex

	for _, expr := range expressions {
		semaphore <- struct{}{} // Acquire semaphore
		wg.Add(1)
		go func(expr string) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release semaphore

			val, err := evaluateExpression(expr, variables)
			if err != nil {
				fmt.Println("Error evaluating expression:", err)
				return
			}

			mu.Lock()
			results[expr] = val
			mu.Unlock()
		}(expr)
	}

	wg.Wait()

	// Print the results
	for expr, result := range results {
		fmt.Printf("Result of %s: %s\n", expr, result)
	}
}

func evaluateExpression(expr string, variables map[string]int) (string, error) {
	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return "", err
	}

	parameters := make(map[string]interface{})
	for key, value := range variables {
		parameters[key] = value
	}

	result, err := expression.Evaluate(parameters)
	if err != nil {
		return "", err
	}

	return strconv.FormatFloat(result.(float64), 'f', -1, 64), nil
}