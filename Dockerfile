FROM golang:1.16.5-alpine3.14
WORKDIR /app
ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download
ADD . /app
RUN go build -o lets-fight .

CMD ["./lets-fight"]