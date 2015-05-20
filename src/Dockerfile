FROM golang

RUN go get github.com/gin-gonic/gin
RUN go get github.com/franela/goreq

ENTRYPOINT go run /go/src/github.com/kickevent/event-api/main.go

VOLUME /go/src/github.com/kickevent/event-api/

EXPOSE 8080
