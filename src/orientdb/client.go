package orientdb

import (
    "fmt"
    "encoding/base64"
    "github.com/franela/goreq"
)

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
