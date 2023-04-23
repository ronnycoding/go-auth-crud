FROM golang:1.17-alpine

ENV GO111MODULE=on

WORKDIR /go/src/github.com/ronnycoding/go-auth-crud

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["/go/src/github.com/ronnycoding/go-auth-crud/main"]