# Streaming Service Catalog API

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat&logo=postgresql&logoColor=white)
![Chi Router](https://img.shields.io/badge/Chi%20Router-4B3263?style=flat&logo=go&logoColor=white)
![JSON API](https://img.shields.io/badge/JSON-000000?style=flat&logo=json&logoColor=white)


## Tech Stack

- Go (Golang)
- Chi or Gorilla Mux (HTTP router)
- PostgreSQL (or SQLite for development)
- JSON API responses
- Optional: Swagger docs, Docker, unit tests

## Project Structure
```
streamflix-api/
├── cmd/
│ └── main.go // Entry point
├── handlers/
│ └── catalog.go // Route handlers
├── models/
│ └── media.go // Media struct and DB logic
├── db/
│ └── connection.go // Database setup
├── routes/
│ └── router.go // Route definitions
├── middleware/
│ └── logger.go // Optional logging
├── .env // Environment config
├── go.mod / go.sum // Go modules
└── README.md
```


## API Endpoints

| Method | Route          | Description           |
|--------|----------------|-----------------------|
| GET    | `/media`       | List all media        |
| GET    | `/media/:id`   | Get a single item     |
| POST   | `/media`       | Add a new item        |
| PUT    | `/media/:id`   | Update an item        |
| DELETE | `/media/:id`   | Delete an item        |

## Key Features

- REST-style routing
- Database queries with `database/sql`
- JSON request/response handling
- Modular code organization

## Notes

This project is inspired by the backend structure of streaming platforms. It is meant to demonstrate skills relevant to those particular tech environments.

