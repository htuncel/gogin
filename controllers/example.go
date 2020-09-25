package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"main/utils"
)

type ExampleController struct{}

type ExampleService interface {
	tokenHandler(c *gin.Context)
}

func NewExampleController() *ExampleController {
	return new(ExampleController)
}

// GET /token
// Get example token
func (e *ExampleController) TokenHandler(c *gin.Context) {
	token, errToken := utils.GenerateToken()
	claims, _ := utils.VerifyToken(token)
	host := utils.GetClientIP(c)
	token2 := utils.GetToken(c)

	if errToken != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errToken.Error()})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"data":   token,
			"claims": claims,
			"host":   host,
			"token2": token2,
		})
	}
}

func (e *ExampleController) MultipleFileUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload"]

	for _, file := range files {
		// TODO could not fix path to /files/filename
		filename := filepath.Base(file.Filename)
		if errSaveFiles := c.SaveUploadedFile(file, filename); errSaveFiles != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", errSaveFiles.Error()))
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
