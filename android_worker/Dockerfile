FROM golang:1.22.2 AS builder
WORKDIR /usr/src/app

# Install CA certificates
RUN apt-get update && apt-get install -y ca-certificates

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o android_worker ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/android_worker ./android_worker
CMD ["./android_worker"]