FROM golang:latest as builder

WORKDIR /go/src/GolangMicroservice/consignment-micro-cli

COPY . .

# pull in dependencies with go dep
RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init && dep ensure


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY  --from=builder /go/src/GolangMicroservice/consignment-micro-cli/consignment-micro-cli .

COPY consignment.json /app/consignment.json

CMD ["./consignment-micro-cli"]
