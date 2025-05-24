# Go Playground

A simple Go project for experimenting with Go programming concepts and features.

## Getting Started

### Prerequisites

- Go 1.24.3 or later

### Installation

1. Clone the repository
   ```
   git clone https://github.com/nunorocorreia/Go-Playground.git
   cd Go-Playground
   ```

2. Build the project
   ```
   go build ./cmd/main.go
   ```

3. Run the project
   ```
   go run ./cmd/main.go
   ```

## Running Tests

Run all tests in the project:
```
go test ./...
```

Run tests with coverage:
```
go test -cover ./...
```

Generate a coverage report:
```
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Project Structure

- `cmd/` - Contains the application entry points
  - `main.go` - Main application entry point

## Module Information

This project uses the Go module name `hello`.