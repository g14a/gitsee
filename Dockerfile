FROM golang:1.14.4-alpine3.12 AS builder
WORKDIR /app
COPY go.sum go.mod .env ./
RUN go mod download
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
FROM scratch
COPY .env /app/
COPY --from=builder /app/main /app/main
CMD ["/app/main"]