package main

type RequestBody struct {
	Exp    map[string]string `json:"exp"`
	Config map[string]int    `json:"config"`
}

func Sub(a, b int) int {
    return a - b
}

func Add(a, b int) int {
    return a + b
}


