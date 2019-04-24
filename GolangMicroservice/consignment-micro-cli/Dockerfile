FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app
ADD consignment.json /app/consignment.json
ADD consignment-micro-cli /app/consignment-micro-cli

CMD ["./consignment-micro-cli"]