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

type Event struct {
    Event string `json:"event" binding:"required"`
    Name  string `json:"event1" binding:"required"`
}

func (event *Event) get(c *gin.Context) {
    client := Client{
        Uri: "http://orientdb:2480",
        Username: "root",
        Password: "0r13ntDB",
    }
    res := client.Request("GET", "database/kickevent")
    fmt.Println(res)
    c.JSON(201, res)
}

func (event *Event) post(c *gin.Context) {
    c.JSON(201, "test")
}

func (event *Event) delete(c *gin.Context) {
    c.JSON(201, "test")
}

// Client OrientDB à sortir dans un package différent
func base64Encode(str string) string {
    return base64.StdEncoding.EncodeToString([]byte(str))
}

type Client struct {
    Uri      string
    Username string
    Password string
}

func (c *Client) Request(method string, uri string) (string) {
    req := goreq.Request{
        Method: method,
        Uri: fmt.Sprintf("%s/%s", c.Uri, uri),
    }
    req.AddHeader("Accept-Encoding", "gzip,deflate")
    req.AddHeader("Authorization", fmt.Sprintf("Basic %s", base64Encode(fmt.Sprintf("%s:%s", c.Username, c.Password))))
    res, _ := req.Do()
    body, _ := res.Body.ToString()

    return body
}
