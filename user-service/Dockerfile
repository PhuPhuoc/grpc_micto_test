FROM golang:1.23.2 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-services .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/user-services .

EXPOSE 50051
CMD ["./user-services"]