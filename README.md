# ecommerce-crud-gwh ✅

A small, opinionated example of a simple e-commerce CRUD API written in Go. The
project demonstrates using the standard library `net/http`, simple middleware
chaining, and an in-memory product store (no database). It is intended as a
learning/example project.

---

## 🔧 Features

- List products (GET /products)
- Get single product by ID (GET /products/{id})
- Create a product (POST /products)
- Middleware for logging and CORS
- Simple in-memory data store pre-seeded with sample products
- Placeholder handlers for update (PUT) and delete (DELETE)

---

## Requirements

- Go 1.25+ (see `go.mod`)

---

## Quick start

1. Clone the repository

```bash
git clone <repo-url>
cd ecommerce-crud-gwh
```

2. Run the server

```bash
go run .
```

The server starts on port 3000 (http://localhost:3000).

---

## API Endpoints 📡

### GET /products

Returns a JSON array of products.

Example:

```bash
curl http://localhost:3000/products
```

Response (200):

```json
[
  {
    "id": 1,
    "title": "Orange",
    "description": "Orange is red. And I love oranges",
    "price": 100.0,
    "img_url": "https://..."
  },
  ...
]
```

### GET /products/{id}

Returns a single product by ID.

Example:

```bash
curl http://localhost:3000/products/1
```

### POST /products

Create a new product. Send a JSON body with the following fields:

- `title` (string)
- `description` (string)
- `price` (number)
- `img_url` (string)

Example:

```bash
curl -X POST http://localhost:3000/products \
  -H "Content-Type: application/json" \
  -d '{"title":"Pear","description":"Fresh pear","price":90,"img_url":"https://..."}'
```

Response (201): created product object.

### PUT /products/{id}

- Not implemented yet (handler present as TODO).

### DELETE /products/{id}

- Not implemented yet (handler present as TODO).

---

## Project structure 🔍

```
./
├─ cmd/                # application bootstrap and route registration
├─ database/           # in-memory product model & seed data
├─ handlers/           # route handlers (create/get/update/delete)
├─ middleware/         # middleware helpers (CORS, logging)
├─ utils/              # helper utilities (SendData)
├─ main.go             # entrypoint
├─ go.mod
```

---

## Notes & TODOs ✨

- Data is stored in memory (`database.ProductList`). Restarting the server will
  reset data.
- `UpdateProduct` and `DeleteProduct` handlers are TODOs and need
  implementation.
- Path parameter extraction (e.g., `/products/{id}`) is used in handlers —
  review request routing if you change the router implementation.

---

## Contributing

Contributions are welcome — open an issue or PR. If you add features, please
include tests and update this README.

---

## License

MIT — see LICENSE (add one if you want to publish this project)
