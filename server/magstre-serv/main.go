package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	controllers "github.com/Voltzz-stv/go-movies/server/magstre-serv/controller"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/movies", controllers.GetMovies())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("failed to start server", err)
	}
}
