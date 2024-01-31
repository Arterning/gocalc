package main

import (
	"fmt"
	"strings"
)

func main() {
	expressions := map[string]string{
		"f": "9+c",
		"c": "32+3",
		"a": "c+3",
		"b": "a/2",
	}

	dependencyGraph := make(map[string][]string)

	for target, expr := range expressions {
		deps := extractDependencies(expr, expressions)
		dependencyGraph[target] = deps
	}

	fmt.Printf("Dependency Graph: %v\n", dependencyGraph)
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