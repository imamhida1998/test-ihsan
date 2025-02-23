
FROM golang:1.23-alpine


ENV GO111MODULE=on


WORKDIR /app


COPY go.mod go.sum ./

# Unduh dependencies
RUN go mod download

COPY . .



RUN go build -o main .

EXPOSE 8080

# Perintah untuk menjalankan aplikasi
CMD ["./main"]