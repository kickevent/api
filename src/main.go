package main

import (
    "fmt"
    "encoding/base64"
    "github.com/gin-gonic/gin"
    "github.com/franela/goreq"
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

func base64Encode(str string) string {
    return base64.StdEncoding.EncodeToString([]byte(str))
}

type Event struct {
    Event string `json:"event" binding:"required"`
    Name  string `json:"event1" binding:"required"`
}

func (event *Event) get(c *gin.Context) {
    // Connect to the database
    req := goreq.Request{
        Uri: "http://orientdb:2480/connect/kickevent",
    }
    req.AddHeader("Accept-Encoding", "gzip,deflate")
    req.AddHeader("Authorization", fmt.Sprintf("Basic %s", base64Encode("root:0r13ntDB")))
    res, err := req.Do()
    if err != nil {
        fmt.Println(err.Error())
    }
    body, _ := res.Body.ToString()
    fmt.Println(body)

    // Getting database informations
    req = goreq.Request{
        Uri: "http://orientdb:2480/database/kickevent",
    }
    req.AddHeader("Accept-Encoding", "gzip,deflate")
    req.AddHeader("Authorization", fmt.Sprintf("Basic %s", base64Encode("root:0r13ntDB")))
    res, err = req.Do()
    if err != nil {
        fmt.Println(err.Error())
    }
    body, _ = res.Body.ToString()
    fmt.Println(body)

    c.JSON(201, body)
}

func (event *Event) post(c *gin.Context) {
    c.JSON(201, "test")
}

func (event *Event) delete(c *gin.Context) {
    c.JSON(201, "test")
}
