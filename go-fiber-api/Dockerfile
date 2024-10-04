# Gunakan image golang untuk membangun aplikasi
FROM golang:1.23-alpine as builder

# Set environment variable
ENV GO111MODULE=on

# Buat direktori kerja di dalam container
WORKDIR /app

# Copy semua file go.mod dan go.sum ke dalam container
COPY go.mod go.sum ./

# Download dependensi yang dibutuhkan
RUN go mod download

# Copy semua kode aplikasi ke dalam container
COPY . .



# Build aplikasi
RUN go build -o main .

# Gunakan image minimal untuk runtime
FROM alpine:latest

# Set direktori kerja di dalam container
WORKDIR /root/

# Copy aplikasi yang telah di-build dari stage builder
COPY --from=builder /app/main .

# Copy file .env ke dalam container
COPY .env .

# Set port aplikasi yang digunakan
EXPOSE 3000

# Perintah untuk menjalankan aplikasi
CMD ["./main"]
