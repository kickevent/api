package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    event := &Event{}

    v1 := router.Group("/v1")
    {
        v1.GET("/event", event.get)
        v1.POST("/event", event.post)
        v1.DELETE("/event", event.delete)
    }

    router.Run(":8080")
}


type Event struct {
    Event string `json:"event" binding:"required"`
    Name  string `json:"event1" binding:"required"`
}

func (event *Event) get(c *gin.Context) {

    c.JSON(201, event)
}

func (event *Event) post(c *gin.Context) {
    c.JSON(201, "test")
}

func (event *Event) delete(c *gin.Context) {
    c.JSON(201, "test")
}
