package recipe

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRecipe(ctx *gin.Context) {
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
