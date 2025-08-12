# Estágio 1: Compilação
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /meu-servidor .

# Estágio 2: Execução
FROM scratch
COPY --from=builder /meu-servidor /meu-servidor
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080
CMD ["/meu-servidor"]
