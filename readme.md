# 🐹 go-rest-api

> REST API built with Go covering code structure, authentication, middleware and package management.

![Go](https://img.shields.io/badge/Go_1.26-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Chi](https://img.shields.io/badge/Chi_Router-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

---

## 📚 What was practiced

- **Golang Code Structure** — packages and modules organization (`cmd/`, `api/`, `internal/`)
- **API Authentication** — JWT-based authentication
- **Middleware** — route protection with Chi middleware
- **Package Management** — dependency management with Go modules

---

## 🗂️ Project Structure
```
go-rest-api/
├── cmd/
│   └── api/
│       └── main.go          ← entry point
├── api/                     ← handlers and routes
├── internal/                ← internal business logic
├── go.mod
└── go.sum
```

---

## 🛠️ Technologies

| Technology | Purpose |
|------------|---------|
| [Go 1.26](https://golang.org) | Main language |
| [Chi](https://github.com/go-chi/chi) | Lightweight HTTP router and middleware |
| [Gorilla Schema](https://github.com/gorilla/schema) | Request decoding |
| [Logrus](https://github.com/sirupsen/logrus) | Structured logging |

---

## ⚙️ How to run
```bash
# Clone the repository
git clone https://github.com/DevLucasHenrique/go-rest-api.git
cd go-rest-api

# Install dependencies
go mod tidy

# Run the server
go run cmd/api/main.go
```

---

## 👨‍💻 Author

[![GitHub](https://img.shields.io/badge/GitHub-DevLucasHenrique-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/DevLucasHenrique)

Made with 💙 by **Lucas Henrique**