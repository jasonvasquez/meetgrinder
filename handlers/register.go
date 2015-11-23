package handlers

import (
	"github.com/jasonvasquez/meetgrinder-api/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jasonvasquez/meetgrinder-api/handlers/events"
)

func Register(r *gin.RouterGroup) {

	r.GET("/events", events.Index)
	r.POST("/events", events.Create)
	r.GET("/events/:id", events.Show)

}
