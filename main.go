package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	engine "go-expression-engine/lib"
)


func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})


    router.POST("/evaluate", func(c *gin.Context) {
		var requestBody RequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(400, gin.H{"error": "Failed to parse request body as JSON"})
			return
		}

		// Print the parsed JSON body
		fmt.Printf("Received JSON body: %+v\n", requestBody)

		result, err := engine.CalcExpressions(requestBody.Exp)

		//response result 
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
			return
		}

		// Send the result as a JSON string to the client
		c.JSON(http.StatusOK, gin.H{"message": "success", "result": result})
	})

	router.Run(":8080")
}