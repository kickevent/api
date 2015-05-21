package main

import (
    "fmt"
    "github.com/kickevent/event-api/orientdb"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    event := &Event{}

    v1 := router.Group("/v1")
    {
        v1.GET("/database/create", event.createDatabase)
    }

    router.Run(":8080")
}

type Event struct {
    Event string `json:"event" binding:"required"`
    Name  string `json:"event1" binding:"required"`
}

func (event *Event) createDatabase(c *gin.Context) {
    client := orientdb.Client{
        Uri: "http://orientdb:2480",
        Username: "root",
        Password: "0r13ntDB",
    }
    res := client.Request("POST", "database/kickevent/plocal")
    res = client.Request("POST", "class/kickevent/event")
    res = client.Request("POST", "property/kickevent/event/name/STRING")
    res = client.Request("POST", "property/kickevent/event/description/STRING")
    fmt.Println(res)
    c.JSON(201, res)
}
