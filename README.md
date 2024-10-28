# Receipt Processor

## Overview
Receipt Processor is a RESTful API built with the Gin framework in Go. It allows users to submit receipts for processing and calculate reward points based on specific criteria. This project demonstrates a practical example of data processing and scoring logic in Go, following clean code principles.

## Features
- Submit a receipt for processing and receive an ID.
- Retrieve reward points awarded to a receipt based on custom rules.
- Input validation, error handling, and response formatting using the Gin framework.

## Tech Stack
- **Language**: Go
- **Framework**: Gin
- **UUID Generation**: Google UUID library

  ## Setup Instructions

### Prerequisites
- [Go](https://golang.org/dl/) (1.16 or later)
- [Gin](https://github.com/gin-gonic/gin) framework
- [Google UUID](https://pkg.go.dev/github.com/google/uuid) library

### Installing Dependencies
1. Clone the repository:
    ```bash
    git clone https://github.com/marehman0123/receipt-processor-challenge.git
    cd receipt-processor-challenge
    ```
2. Install required Go packages:
    ```bash
    go mod tidy
    ```

### Running the Server
Run the server on `localhost:8080` with:
```bash
go run main.go
