package main

import (
    "fmt"
    "github.com/kickevent/event-api/orientdb"
    "github.com/gin-gonic/gin"
)

const DatabaseName = "kickevent"
const DatabaseUser = "root"
const DatabasePassword = "0r13ntDB"
const DatabaseHost = "http://orientdb:2480"

func main() {
    router := gin.Default()

    v1 := router.Group("/v1")
    {
        v1.POST("/database", createDatabase)
        v1.POST("/event", createEvent)
        v1.GET("/event/:id", getEvent)
        v1.PUT("/event/:id", updateEvent)
        v1.DELETE("/event/:id", deleteEvent)
    }

    router.Run(":8080")
}

type Event struct {
    Name string `json:"name" binding:"required"`
    Description  string `json:"description" binding:"required"`
}

func getOrientdbClient() (orientdb.Client) {
    return orientdb.Client{
        Uri: DatabaseHost,
        Username: DatabaseUser,
        Password: DatabasePassword,
    }
}

func createDatabase(c *gin.Context) {
    client := getOrientdbClient()
    _ = client.Request("POST", fmt.Sprintf("database/%s/plocal", DatabaseName))
    _ = client.Request("POST", fmt.Sprintf("class/%s/event", DatabaseName))
    _ = client.Request("POST", fmt.Sprintf("property/%s/event/name/STRING", DatabaseName))
    _ = client.Request("POST", fmt.Sprintf("property/%s/event/description/STRING", DatabaseName))
    c.JSON(200, "OK")
}

func createEvent(c *gin.Context) {
    var event Event
    c.Bind(&event)
    body := map[string]interface{}{
        "@class": "event",
        "name":  event.Name,
        "description": event.Description,
    }
    client := getOrientdbClient()
    res := client.RequestWithBody("POST", fmt.Sprintf("document/%s", DatabaseName), body)
    c.JSON(200, res)
}

func updateEvent(c *gin.Context) {
    var event Event
    c.Bind(&event)
    body := map[string]interface{}{
        "@class": "event",
        "name":  event.Name,
        "description": event.Description,
    }
    client := getOrientdbClient()
    res := client.RequestWithBody("PUT", fmt.Sprintf("document/%s/%s", DatabaseName, c.Params.ByName("id")), body)
    c.JSON(200, res)
}

func deleteEvent(c *gin.Context) {
    client := getOrientdbClient()
    res := client.Request("DELETE", fmt.Sprintf("document/%s/%s", DatabaseName, c.Params.ByName("id")))
    c.JSON(200, res)
}

func getEvent(c *gin.Context) {
    client := getOrientdbClient()
    res := client.Request("GET", fmt.Sprintf("document/%s/%s", DatabaseName, c.Params.ByName("id")))
    c.JSON(200, res)
}