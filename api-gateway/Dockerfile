# --- Estágio 1: Build ---
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN chmod +x wait-for-it.sh
CMD ["./wait-for-it.sh", "rabbitmq:5672", "--", "./main"]

# Compila a aplicação, e as flags garantem que o binário seja estático, bom para imgs pequenas
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# --- Estágio 2: Produção ---
# Começa uma imagem "scratch", que é a menor imagem possível (vazia, mais segura)
FROM scratch

WORKDIR /app

# Copia APENAS o arquivo binário compilado do estágio de build
# Assim, nenhum código-fonte ou ferramenta de build vai para a imagem final
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]