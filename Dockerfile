FROM docker.io/library/golang:1.21-alpine3.19

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

CMD ["/main"]