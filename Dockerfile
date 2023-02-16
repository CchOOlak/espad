FROM golang:1.19-alpine

WORKDIR /go/src/app
COPY . .

RUN go get ./...

COPY ./cmd/ .

RUN go build -o app .

ENTRYPOINT ["./app"]
