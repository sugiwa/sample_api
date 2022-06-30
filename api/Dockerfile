FROM golang:1.18-alpine

RUN apk add --update && apk add git

WORKDIR /go/src/app

CMD [ "go", "run", "main.go" ]

EXPOSE 8080
 