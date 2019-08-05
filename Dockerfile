FROM alpine:latest

RUN apk add --no-cache libc6-compat
RUN mkdir /app
WORKDIR /app
ADD consignment-service /app/consignment-service

CMD ["./consignment-service"]