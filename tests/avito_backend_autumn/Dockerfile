FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine

WORKDIR /build

EXPOSE 8080
COPY --from=builder /build/main /build/main

RUN ls

CMD ["./main"]
