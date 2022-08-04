#build stage
FROM golang:1.18-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]