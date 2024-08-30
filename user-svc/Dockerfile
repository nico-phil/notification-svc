FROM golang:1.22.2 AS builder

WORKDIR /usr/src/app

# Install necessary packages like wget and ca-certificates
RUN apt-get update && apt-get install -y wget ca-certificates

# Verify the installation of wget and ca-certificates
RUN wget --version && update-ca-certificates

# Download and install the migrate tool
RUN wget --no-check-certificate -O /usr/local/bin/migrate https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.deb \
    && chmod +x /usr/local/bin/migrate

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-svc ./cmd/main.go



FROM scratch
COPY --from=builder /usr/src/app/user-svc ./user-svc

CMD ["migrate", "-path", "migrations", "-database", "postgres://postgres:postgres@postgres/users?sslmode=disable", "up", "./user-svc"]
CMD ["./user-svc"]