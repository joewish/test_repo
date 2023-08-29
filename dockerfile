FROM golang:latest
COPY main.go .
RUN go run main.go