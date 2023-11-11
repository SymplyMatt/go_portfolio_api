package main

import (
	"os"

	"github.com/SymplyMatt/go_portfolio_api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/project/create", routes.AddProject)
	router.GET("/projects", routes.GetProjects)
	router.PUT("/project/update/:id", routes.UpdateProject)
	router.DELETE("/project/delete/:id", routes.DeleteProject)

	router.Run(":" + port)
}
