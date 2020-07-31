FROM golang:1.14.4-alpine3.12 AS builder
RUN mkdir /app
WORKDIR /app
COPY go.sum .
COPY go.mod .
RUN go mod download
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
FROM scratch
COPY --from=builder /app/main /app/main
CMD ["/app/main"]