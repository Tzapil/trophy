package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tzapil/trophy/game"
	"github.com/tzapil/trophy/user"
)

func serve() {
	// creating of new router
	r := gin.Default()

	// make all handlers v1 api version
	v1 := r.Group("/api/v1")

	user.UserRegister(v1)
	game.GameRegister(v1)
	// collections.CollectionsRegister(v1)

	// new handler /ping
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// run default listen and serve on 0.0.0.0:8080
	r.Run()
}

func main() {
	// log.Parse()

	serve()
}
