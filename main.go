package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type recipe struct {
	Title       string
	Ingredients []string
	Steps       []string
}

func getRecipe(ctx *gin.Context) {
	recipe := recipe{
		Title: "Asian pork rice",
		Ingredients: []string{
			"pork",
			"rice",
			"onion",
			"garlic",
			"soy sauce",
			"sugar",
			"honey",
			"water",
			"mirin",
		},
		Steps: []string{
			"Fry the pork",
			"cut into bite sized pieces",
			"Put in the onion and garlic",
			"cook",
		},
	}

	out, err := json.Marshal(recipe)
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(out))

	ctx.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()

	// Create API route group
	api := router.Group("/api")
	{
		// Add /hello GET route to router and define route handler function
		api.GET("/recipe", getRecipe)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	router.Run(":8090")
}
