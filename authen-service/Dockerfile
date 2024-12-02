FROM golang:1.23.2 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o authen-service .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/authen-service .

EXPOSE 50051
CMD ["./authen-service"]
