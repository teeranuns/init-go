# init-go
start learning golang 

## Project Structure

```
📂 src/
├─ 📄 main.go
📂 assets/
├─ 📂 logs/
📂 configs/
├─ 📄 configs.go
📂 domain/
├─ 📂 servers/
│  ├─ 📄 server.go
│  ├─ 📄 handler.go
├─ 📂 entities/
│  ├─ 📄 users.go
│  ├─ 📄 response.go
├─ 📂 users/
│  ├─ 📂 controllers/
│  │  ├─ 📄 users_controllers.go
│  ├─ 📂 usecases/
│  │  ├─ 📄 users_usecases.go
│  ├─ 📂 repositories/
│  │  ├─ 📄 users_repositories.go
│  ├─ 📂 models/
│  │  ├─ 📄 users_response.go
📂 pkg/
├─ 📂 databases/
│  ├─ 📂 migrations/
│  ├─ 📄 postgresql.go
├─ 📂 middlewares/
├─ 📂 utils/
│  ├─📄 connection_url_builder.go
📂 tests/
├─ 📂 users/
│  ├─ 📄 users_test.go
📄 .env
```

## Getting Started
- air

### Prerequisites

- Go 1.22.4
- MySQL
- [Echo Framework](https://echo.labstack.com/)