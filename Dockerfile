FROM golang:1.26

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o tradebot

CMD ["./tradebot"]
