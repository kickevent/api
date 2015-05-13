FROM golang

COPY ./ /app

RUN go get github.com/gin-gonic/gin
RUN go get github.com/franela/goreq

ENTRYPOINT go run /app/main.go

VOLUME /app

EXPOSE 8080
