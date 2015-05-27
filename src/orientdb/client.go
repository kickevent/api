package orientdb

import (
    "fmt"
    "encoding/json"
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

type Body interface {
}

func (c *Client) RequestWithBody(method string, uri string, body Body) (Body) {
    req := goreq.Request{
        Method: method,
        Uri: fmt.Sprintf("%s/%s", c.Uri, uri),
        Body: body,
    }
    req.AddHeader("Accept-Encoding", "gzip,deflate")
    req.AddHeader("Authorization", fmt.Sprintf("Basic %s", base64Encode(fmt.Sprintf("%s:%s", c.Username, c.Password))))
    res, _ := req.Do()
    resbody, _ := res.Body.ToString()
    src_json := []byte(resbody)
    var jsonresbody interface{}
    _ = json.Unmarshal(src_json, &jsonresbody)
    
    return jsonresbody
}

func (c *Client) Request(method string, uri string) (Body) {
    return c.RequestWithBody(method, uri, nil)
}