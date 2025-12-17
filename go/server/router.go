package server

import (
	"net/http"
	"recipe-app/go/recipe"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	{
		api := router.Group("/api")
		api.GET("/recipe", recipe.GetRecipe)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
