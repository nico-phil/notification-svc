FROM golang:1.22.2 AS builder
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o email_worker ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/email_worker ./email_worker
CMD ["./email_worker"]