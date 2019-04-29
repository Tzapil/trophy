package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GameRegister(router *gin.RouterGroup) {
	router.GET("/game/:name", GameRetrieve)
}

func GameRetrieve(c *gin.Context) {
	name := c.Param("name")
	result := GetGame(name)

	c.JSON(http.StatusOK, gin.H{"game": result})
}
