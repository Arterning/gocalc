package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Variable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ExpressionRequest struct {
	Exp []Variable `json:"exp"`
	J   int        `json:"j"`
}

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})


    router.POST("/evaluate", func(c *gin.Context) {
		var req ExpressionRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var result []Variable
		for _, exp := range req.Exp {
			value, err := evaluateExpression(exp.Value, req.Exp)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			result = append(result, Variable{Name: exp.Name, Value: value})
		}
		c.JSON(http.StatusOK, gin.H{"out": result})
	})

	router.Run(":8080")
}


func evaluateExpression(expr string, variables []Variable) (string, error) {
	for _, v := range variables {
		expr = strings.Replace(expr, v.Name, v.Value, -1)
	}
	val, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		log.Println("Error creating expression")
		return "", err
	}

   
    variablesMap := make(map[string]interface{})
    for _, v := range variables {
		//Cast v to float
		v.Value = strconv.FormatFloat(v.Value, 'f', -1, 64)
        variablesMap[v.Name] = v
    }


	result, err := val.Evaluate(variablesMap)
    
	if err != nil {
		log.Println("Error evaluating expression")
		return "", err
	}
	return strconv.FormatFloat(result.(float64), 'f', -1, 64), nil
}
