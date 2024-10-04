go-fiber-app/
│
├── cmd/
│   └── main.go           # Entry point aplikasi
│
├── config/
│   └── config.go         # Konfigurasi aplikasi (misal port berapa)
│
├── database/
│   └── database.go         # Konfigurasi database
|
├── internal/
│   ├── controller/         # Controller untuk menangani request HTTP
│   │   └── user_controller.go
│   │
│   ├── models/           # Definisi model untuk database atau objek
│   │   └── user.go
│   │
│   ├── repositories/     # Interaksi dengan "database" atau storage
│   │   └── user_repository.go
│   │
│   ├── services/         # Logika bisnis utama
│   │   └── user_service.go
│   │
│   └── routes/           # Definisi rute API
│       └── routes.go
│
├── tests/
│   └── user_test.go      # Unit test
│
├── go.mod                # Go module file
└── go.sum                # Go dependencies
└── .env                  # Env

====================================================

# Running 
go run cmd/main.go

# Testing
go test tests/user_test.go

# Build Image
1. create Dockerfile and configure
2. build image Golang

# Run Go
docker run -p 3000:3000 go-fiber-api:v.0.1