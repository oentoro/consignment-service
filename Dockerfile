FROM alpine:latest

RUN apk add --no-cache \
        libc6-compat
WORKDIR /app

ADD consignment-service /app/consignment-service

CMD ["./consignment-service"]