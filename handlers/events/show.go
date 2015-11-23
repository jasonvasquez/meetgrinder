package events

import (
	"github.com/jasonvasquez/meetgrinder-api/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jasonvasquez/meetgrinder-api/model"
)

func Show(c *gin.Context) {
	var eventId = c.Param("id")

	event := model.FindEventById(eventId)
	c.JSON(200, event)
}
