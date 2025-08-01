# Estágio 1: Construção
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Estágio 2: Produção
FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]