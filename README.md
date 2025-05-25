# Go Shopping Cart API
A RESTful API for Shopping Cart application with Go, Postgresql.

It is a simple RESTful API with Go using **gin-gonic/gin** (a HTTP web framework) and **gorm** (An ORM library).

## Installation & Run
#### Clone the repository:
   ```
   git clone git@github.com:ndt-hub/kart-server.git
   cd kart-server
   ```

#### Configuration
Before running API server, you should set the database/api config with yours by creating an `.env` file similar to [.env.example](https://github.com/ndt-hub/kart-server/blob/develop/.env.example)
```
API_KEY=apitest

APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=kart_db
DB_SSL_MODE=disable

ALLOWED_ORIGINS=http://localhost:3000,https://example.com
```

#### Installing dependencies
```bash
go mod tidy
```

#### Database migration and initialize sample data
```bash
go run cmd/migrate/main.go
go run cmd/seed/main.go
```

#### Start Server for development
```bash
go run main.go

# API Endpoint listening at: http://127.0.0.1:8080
```

#### Build
```bash
# Build and Run
go build -o build/kart-server
./build/kart-server

# API Endpoint listening at: http://127.0.0.1:8080
```

## Structure
```
kart-server/
├── cmd/
│   ├── migrate/         // Migration command
│   └── seed/            // Seeding command
├── config               // Database and Environment configuration
├── controller/
│   ├── discount.go      // APIs for Discount model
│   ├── order.go         // APIs for Order model
│   └── product.go       // APIs for Product model
├── model/
│   ├── discount.go      // Discount model
│   ├── order.go         // Order model
│   └── product.go       // Product model
├── db/
│   ├── migrations/      // Database migrations
│   └── seeds/           // Database seeds
├── public/
│   └── images/          // Public images directory
└── main.go              // Main application entry point

```

## API

#### /product
* `GET` : Get all product
* `POST` : Create a new product

#### /product/:productId
* `GET` : Get a product details

#### /discount/:discountCode
* `GET` : Get a discount details

#### /order
* `POST` : Place an order
