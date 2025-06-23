FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build 

FROM alpine:3.17

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/taskd ./

COPY config.yaml ./

EXPOSE 8080

ENTRYPOINT ["./taskd"]