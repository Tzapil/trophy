package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(router *gin.RouterGroup) {
	router.GET("/user/:name/logs", UserLogsRetrieve)
	router.GET("/user/:name", UserRetrieve)
}

func UserLogsRetrieve(c *gin.Context) {
	name := c.Param("name")
	results := GetLogs(name)

	c.JSON(http.StatusOK, gin.H{"items": results, "size": len(results)})
}

func UserRetrieve(c *gin.Context) {
	name := c.Param("name")
	results := GetUser(name)

	c.JSON(http.StatusOK, gin.H{"user": results})
}
