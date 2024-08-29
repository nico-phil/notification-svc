FROM golang:1.22.2 AS builder
WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o notification-svc ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/notification-svc ./notification-svc
CMD ["./notification-svc"]