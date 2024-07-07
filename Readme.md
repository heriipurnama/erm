# ERM API

## Description

This project is a RESTful API application for managing customers, orders, and authentication using Golang, Gin-Gonic, GORM, and MySQL. It provides endpoints for customer management, order management, and user authentication, with support for pagination and JWT-based authentication.

## Technologies

- **Golang** - Programming language used for the API.
- **Gin-Gonic** - Web framework for building the HTTP server.
- **GORM** - ORM for interacting with the MySQL database.
- **MySQL** - Database system for storing data.
- **Docker** - Containerization platform for running the application.
- **JWT** - JSON Web Token for secure authentication.

## Features

- **Customer Management**:
  - Get all customers (with pagination)
  - Get customer details by ID
  - Create a new customer
  - Update customer details
  - Delete a customer
  - Search for customers

- **Order Management**:
  - Get all orders (with pagination)
  - Get order details by ID
  - Create a new order
  - Update order details
  - Delete an order
  - Search for orders

- **Authentication Management**:
  - Register a new user
  - Login and receive a JWT token

## Setup

### 1. Clone the Repository

```bash
git@github.com:heriipurnama/erm.git
cd erm
```

### 2. Install Dependencies

Ensure you have Go installed. Run the following command to install the project dependencies:

```bash
go mod tidy
```

### 3. Create a `.env` File

Create a `.env` file in the root directory with the following content:

```env
DB_DSN="username_db:password_db@tcp(host_db:port_db)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
```

### 4. Run Database Migrations

Run the application to automatically apply database migrations:

```bash
go run main.go
```

### 5. Start the Application

Run the application:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

### 6. Docker Setup

To build and run the application using Docker, use the following commands:

- **Build the Docker image**:

```bash
docker build -t erm-api .
```

- **Run the Docker container**:

```bash
docker run -p 8080:8080 --env-file .env erm-api
```

## API Endpoints

### Authentication

- **Login**: `POST /login`
  - **Request Body**: `{"username": "user", "password": "password"}`
  - **Response**: JWT token

- **Register**: `POST /register`
  - **Request Body**: `{"username": "user", "password": "password"}`

### Customers

- **Get All Customers**: `GET /customers?page=1&limit=10`
- **Get Customer by ID**: `GET /customers/:id`
- **Create Customer**: `POST /customers`
  - **Request Body**: `{"name": "John Doe", "email": "john.doe@example.com", "address": "123 Main St"}`
- **Update Customer**: `PUT /customers/:id`
  - **Request Body**: `{"name": "John Doe Updated", "email": "john.doe.updated@example.com", "address": "456 Another St"}`
- **Delete Customer**: `DELETE /customers/:id`
- **Search Customers**: `GET /customers/search?q=name`

### Orders

- **Get All Orders**: `GET /orders?page=1&limit=10`
- **Get Order by ID**: `GET /orders/:id`
- **Create Order**: `POST /orders`
  - **Request Body**: `{"customer_id": 1, "product": "Widget", "quantity": 10, "price": 19.99}`
- **Update Order**: `PUT /orders/:id`
  - **Request Body**: `{"product": "Widget Updated", "quantity": 15, "price": 29.99}`
- **Delete Order**: `DELETE /orders/:id`
- **Search Orders**: `GET /orders/search?q=product`

### Postman Collection
- **on Dir**: `/docs/postman_collection`

## ERD And Table Database Specifications

### ERD
- **ERD on Dir**: `/docs/ERD`

### Table Database Specifications/db migrations
- **on Dir**: `/db/migrations`

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gin-Gonic](https://github.com/gin-gonic/gin) - HTTP web framework
- [GORM](https://gorm.io/) - ORM library for Golang
- [MySQL](https://www.mysql.com/) - Database system
- [Docker](https://www.docker.com/) - Containerization platform
- [JWT](https://jwt.io/) - Authentication standard
