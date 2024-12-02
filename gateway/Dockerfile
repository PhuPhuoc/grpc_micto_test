FROM golang:1.23.2 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gateway .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/gateway .

EXPOSE 50051
CMD ["./gateway"]