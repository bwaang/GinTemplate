// handlers.user.go

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	// Call the render function with the name of the template to render
	c.JSON(http.StatusOK, gin.H{"status": "Ok"})
}

func versionCheck(c *gin.Context) {
	// Call the render function with the name of the template to render
	c.JSON(http.StatusOK, gin.H{"version": "v0.0.0"})
}
