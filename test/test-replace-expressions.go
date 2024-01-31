package main

import (
	"fmt"
	"strings"
)


func main() {
	expressions := map[string]string{
		"ji":"32+3",
		"f": "9+c",
		"c": "32+3",
		"a": "c+3",
		"b": "a/2",
		"dab":"b+234/3",
		"asset":"234+234/34",
		"shareRatio":"32/23 + 56",
	}

	//handle expressions, replace variables with values
	for key, expr := range expressions {
	    for varName, varValue := range expressions {
	        expr = strings.Replace(expr, varName, "(" + varValue + ")", -1)
	    }
	    expressions[key] = expr
	}

	fmt.Printf("Expressions: %v\n", expressions)
}