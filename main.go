package main

import (
	"github.com/jasonvasquez/meetgrinder-api/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jasonvasquez/meetgrinder-api/handlers"
)

const listenPort = 8081

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.RedirectFixedPath = true

	v1 := r.Group("/api/v1")
	{
		handlers.Register(v1)
	}

	r.Run(":8081")

}
