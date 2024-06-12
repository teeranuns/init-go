# init-go
start learning golang 

## Project Structure

```
project/
├── domain/
│ ├── user.go
│ ├── product.go
│ ├── receive.go
├── dto/
│ ├── user/
│ │ ├── create_user_request.go
│ │ ├── update_user_request.go
│ │ ├── search_user_request.go
│ │ ├── user_response.go
│ ├── product/
│ │ ├── create_product_request.go
│ │ ├── update_product_request.go
│ │ ├── search_product_request.go
│ │ ├── product_response.go
│ ├── receive/
│ │ ├── create_receive_request.go
│ │ ├── update_receive_request.go
│ │ ├── search_receive_request.go
│ │ ├── receive_response.go
├── infrastructure/
│ ├── router.go
├── interfaces/
│ ├── user_controller.go
│ ├── product_controller.go
│ ├── receive_controller.go
├── repository/
│ ├── user_repository.go
│ ├── product_repository.go
│ ├── receive_repository.go
├── usecase/
│ ├── user_usecase.go
│ ├── product_usecase.go
│ ├── receive_usecase.go
├── main.go
```

## Getting Started
- air

### Prerequisites

- Go 1.22.4
- MySQL
- [Echo Framework](https://echo.labstack.com/)