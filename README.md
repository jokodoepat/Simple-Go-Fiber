# Simple Go Fiber

### About
This repository contains a simple CRUD application built with Go Fiber using the MVC architecture. The application uses PostgreSQL as the database and can be run in a container environment such as Docker.

### Structure
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
└── .env                  # Env
