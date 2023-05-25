FROM golang:1.19 AS development
WORKDIR /go/src/github.com/escalopa/table/
COPY . .
RUN go mod download
RUN go install github.com/cespare/reflex@latest
CMD reflex -sr '\.go$' go run ./cmd/main.go

FROM golang:alpine AS builder
WORKDIR /go/src/github.com/escalopa/table/
COPY . .
RUN go build -o /go/bin/table ./cmd

FROM alpine:latest AS production
RUN apk add --no-cache tzdata
COPY --from=builder /go/bin/table /go/bin/table
COPY ./csv /csv
ENTRYPOINT ["/go/bin/table"]
