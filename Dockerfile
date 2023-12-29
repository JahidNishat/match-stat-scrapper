FROM golang:alpine AS builder

WORKDIR /app

ADD go.mod go.mod
ADD go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o myapp

FROM alpine

WORKDIR /app

RUN apk update
RUN apk upgrade
RUN apk add --no-cache chromium

COPY --from=builder /app/myapp /app/myapp
COPY --from=builder /app/config.yaml /app/config.yaml

ENTRYPOINT ["./myapp"]