FROM golang:latest

WORKDIR /

COPY ./ ./

RUN go mod download

RUN go build -o main ./cmd/main.go

CMD ["./main"]
