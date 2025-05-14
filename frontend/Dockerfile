FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o frontend main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/frontend .
COPY . .
EXPOSE 8080
CMD ["./frontend"]