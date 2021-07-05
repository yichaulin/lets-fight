FROM golang:1.16.5-alpine3.14 AS builder
WORKDIR /app
ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download
ADD . /app
RUN go build -o lets-fight .

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/lets-fight .

CMD ["./lets-fight"]