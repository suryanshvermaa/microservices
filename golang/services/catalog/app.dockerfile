FROM golang:1.24-alpine3.22 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

RUN go build -o catalog ./services/catalog/cmd/catalog/main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /build/catalog .

RUN addgroup --system --gid 1001 golang
RUN adduser --system -uid 1001 golang

USER golang
EXPOSE 8080

CMD ["./catalog"]

