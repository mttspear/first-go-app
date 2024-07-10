# Album Management API

This project is a simple RESTful API built with Go and the Gin framework. The API allows for managing a collection of music albums, including operations to list, add, and retrieve albums by ID.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
  - [Example Requests](#example-requests)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

## Project Structure

app/
├── cmd/
│ └── main.go # Main entry point of the application
├── internal/
│ └── album/
│ ├── album.go # Album model and handler functions
│ └── handler.go # Handlers for album routes
├── go.mod # Go module file
├── go.sum # Go dependencies file
└── .gitignore # Git ignore file


## Installation

1. Clone the repository:

```
git clone https://github.com/your-username/your-repo.git
cd your-repo
```


2. Initialize Go modules and install dependencies:

```
go mod tidy
```

## Usage

To start the server, run:

```
go run cmd/main.go
```

By default, the server will run on localhost:8080.
API Endpoints

- GET /albums: Retrieve a list of all albums.
- POST /albums: Add a new album.
- GET /albums/: Retrieve an album by its ID.

Example Requests

- GET /albums

```
curl http://localhost:8080/albums
```
    
- POST /albums

```
curl -X POST http://localhost:8080/albums -d '{"id": "4", "title": "New Album", "artist": "New Artist", "price": 29.99}'
```

- GET /albums/

```
curl http://localhost:8080/albums/1
```

# License

This project is licensed under the MIT License. See the LICENSE file for details.