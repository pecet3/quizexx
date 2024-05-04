# Etap 1: Zbuduj aplikację
FROM golang:latest AS builder

# Ustaw ścieżkę roboczą w kontenerze
WORKDIR /app

# Skopiuj pliki go.mod i go.sum i pobierz zależności
COPY src/go.mod src/go.sum ./
RUN go mod download

# Skopiuj cały kod źródłowy do kontenera
COPY ./src .

# Skompiluj aplikację do binarnego pliku wykonywalnego
RUN CGO_ENABLED=0 GOOS=linux go build -o /quizex

ENV $(cat .env | xargs)

# Etap 2: Przygotuj lekki obraz do uruchomienia
FROM alpine:latest

# Utwórz katalog na pliki 'view' i skopiuj je
WORKDIR /app
COPY --from=builder /app/view ./view

COPY --from=builder /quizex .

EXPOSE 8090

# Uruchom aplikację
CMD ["./quizex"]
