package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"main/controllers"
)

func main() {
	// Log file.
	f, _ := os.Create("server.log")

	// Write the logs to file and console at the same time. remove os.Stdout for only file logging
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.Use(cors.Default())
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	setupRoutes(router)

	errGin := router.Run(":8080")

	if errGin != nil {
		panic("Error Running Gin")
	}
}

func setupRoutes(router *gin.Engine) {
	bookController := controllers.NewBookController()
	exampleController := controllers.NewExampleController()

	v1 := router.Group("/api/v1")
	{
		example := v1.Group("/example")
		{

			example.GET("/token", exampleController.TokenHandler)
			example.POST("/multifile", exampleController.MultipleFileUpload)
		}

		auth := v1.Group("/auth")
		{
			auth.Use(Authenticate())
			auth.GET("/", homeHandler)
		}

		books := v1.Group("/books")
		{
			books.GET("/", bookController.FindBooks)
			books.POST("/", bookController.CreateBook)
			books.GET("/:id", bookController.FindBook)
			books.PATCH("/:id", bookController.UpdateBook)
			books.DELETE("/:id", bookController.DeleteBook)
		}
	}
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
}
