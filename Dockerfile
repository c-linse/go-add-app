FROM golang:1.22-alpine3.19 as builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
