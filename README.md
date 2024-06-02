# Go API

This is a Go application using the Gin framework and GORM ORM to create a RESTful API with CRUD operations to manage users.

## Project Structure

```plaintext
go-api/
├── main.go
├── controllers/
│   └── userController.go
├── models/
│   └── user.go
├── database/
│   └── database.go
└── routes/
    └── routes.go
```

## Prerequisites
Go 1.22 or later
SQLite3

## Initialize the Go module and install dependencies:
```bash
go mod tidy
```

## Running the Application
```bash
go run main.go
```
