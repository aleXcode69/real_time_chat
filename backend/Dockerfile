FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o backend main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/backend .
EXPOSE 50051
CMD ["./backend"]