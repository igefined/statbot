FROM golang:1.19-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bin/app .

FROM alpine:3.16

COPY --from=builder /app/bin/app /app

CMD [ "/app" ]
