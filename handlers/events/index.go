package events

import (
	"github.com/jasonvasquez/meetgrinder-api/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jasonvasquez/meetgrinder-api/model"
)

func Index(c *gin.Context) {
	events := model.FindAllEvents()

	c.JSON(200, events)
}
