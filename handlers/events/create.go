package events

import (
	"fmt"

	"github.com/jasonvasquez/meetgrinder-api/Godeps/_workspace/src/github.com/gin-gonic/gin"

	"github.com/jasonvasquez/meetgrinder-api/model"
)

func Create(c *gin.Context) {
	fmt.Println("Inside the Create method")

	var event model.Event

	if err := c.BindJSON(&event); err != nil {
		fmt.Println("Failed bind:", err)
		panic(err)
	}

	fmt.Println("Bound the event model:", event)
	c.JSON(200, event.Create())
}
