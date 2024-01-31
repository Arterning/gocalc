package lib

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"strings"
	"sync"
)

func Test() {

	var expressions = map[string]string{
		"ji":         "32+3",
		"f":          "9+c",
		"c":          "32+3",
		"a":          "c+3",
		"b":          "a/2",
		"dab":        "b+234/3",
		"asset":      "234+234/34",
		"shareRatio": "32/23 + 56",
	}

	results, err := CalcExpressions(expressions)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Results: %v\n", results)
	}

}

func CalcExpressions(expressions map[string]string) (map[string]interface{}, error) {

	evaluatedValues := make(map[string]interface{})
	var wg sync.WaitGroup
	var mu sync.Mutex

	//build dependency graph
	dependencyGraph := make(map[string][]string)

	for target, expr := range expressions {
		deps := extractDependencies(expr, expressions)
		dependencyGraph[target] = deps
	}

	fmt.Printf("Dependency Graph: %v\n", dependencyGraph)

	for expr := range expressions {
		// if expr has no dependencys, evaluate it immediately
		if dependencyGraph[expr] == nil || len(dependencyGraph[expr]) == 0 {
			wg.Add(1)
			go func(e string) {
				defer wg.Done()
				result, err := evaluateExpression(expressions[e])
				if err != nil {
					fmt.Printf("Error evaluating expression %s: %v\n", e, err)
					return
				}
				mu.Lock()
				evaluatedValues[e] = result
				mu.Unlock()
			}(expr)
		} else {
			result, err := evaluateExpressionWithDependencies(expr, expressions[expr], evaluatedValues, &mu,
				expressions, dependencyGraph)
			if err != nil {
				fmt.Printf("Error evaluating expression %s: %v\n", expr, err)
				return nil, err
			}
			mu.Lock()
			evaluatedValues[expr] = result
			mu.Unlock()
		}
	}

	wg.Wait()

	for key, value := range evaluatedValues {
		fmt.Printf("Result for %s: %v\n", key, value)
	}

	return evaluatedValues, nil
}

func evaluateExpression(expr string) (interface{}, error) {

	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return nil, err
	}

	parameters := make(map[string]interface{})
	result, err := expression.Evaluate(parameters)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func evaluateExpressionWithDependencies(exprID,
	expr string,
	evaluated map[string]interface{},
	mu *sync.Mutex,
	expressions map[string]string,
	dependencyGraph map[string][]string) (interface{}, error) {
	depValues := make(map[string]interface{})
	for _, dep := range dependencyGraph[exprID] {
		if evaluated[dep] == nil {
			value, _ := evaluateExpressionWithDependencies(dep, expressions[dep], evaluated, mu,
				expressions, dependencyGraph)
			mu.Lock()
			evaluated[dep] = value
			mu.Unlock()
		}
		depValues[dep] = evaluated[dep]
	}

	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return nil, err
	}

	parameters := make(map[string]interface{})
	//put dependency values
	for key, value := range depValues {
		parameters[key] = value
	}
	result, err := expression.Evaluate(parameters)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func extractDependencies(expr string, expressions map[string]string) []string {
	var dependencies []string

	for key := range expressions {
		if strings.Contains(expr, key) {
			dependencies = append(dependencies, key)
		}
	}

	return dependencies
}
