# Task API Gateway

A lightweight task management RESTful API written in Go, using Cobra CLI and `http.ServeMux`.

---

## Usage: Makefile Commands

Build and test locally:

```bash
# Build for local platform (binary output in ./bin)
make build-local

# Run local unit tests (includes auth checks)
make test
```

You can run the binary with custom port and log level:

```bash
./bin/taskd --port 8081 --log-level debug
```

Docker build and run locally:
```bash
# Build and run docker
make docker-run
```

---

## 🔐 Authorization Header Required

All API requests must include the following HTTP header:

```
Authorization: Bearer secret-token
```

Otherwise, the server will respond with `401 Unauthorized` or `403 Forbidden`.

---

## 📚 Task API Endpoints

| Method | Endpoint       | Description            |
|--------|----------------|------------------------|
| GET    | `/tasks`       | Retrieve all tasks     |
| POST   | `/tasks`       | Create a new task      |
| PUT    | `/tasks/{id}`  | Update an existing task |
| DELETE | `/tasks/{id}`  | Delete a task by ID    |

---

## ✅ Task JSON Example

```json
{
  "name": "Task1",
  "status": 0
}
```

- `status`: Task status
  - `0`: InCompleted
  - `1`: Completed