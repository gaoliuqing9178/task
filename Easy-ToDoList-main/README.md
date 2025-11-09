# To-Do List API

This is a simple and robust To-Do List API server built with Go (Golang) and the Gin framework. It provides endpoints for user management (registration, login) and full CRUD (Create, Read, Update, Delete) operations for to-do items, secured by JWT authentication.

## Features

- **User Authentication**: Secure registration and login system using JWT (JSON Web Tokens).
- **CRUD for To-Dos**: Full create, read, update, and delete functionality for to-do items.
- **RESTful API**: Clean and idiomatic RESTful API design.
- **Structured Logging**: Detailed logging for monitoring and debugging.
- **Configuration Management**: External configuration using a `config.toml` file.
- **API Documentation**: Integrated Swagger for easy API exploration and testing.
- **Layered Architecture**: Well-organized project structure (`cmd`, `internal`, `pkg`) for maintainability.

## Getting Started

Follow these instructions to get the project up and running on your local machine.

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or higher recommended)
- A running database instance (the project is configured for a local file-based database like SQLite).

### Installation & Setup

1.  **Clone the repository:**
    ```sh
    git clone <your-repository-url>
    cd todolist
    ```

2.  **Install dependencies:**
    The project uses Go Modules. Dependencies will be downloaded automatically when you build or run the project. You can also download them manually:
    ```sh
    go mod tidy
    ```

3.  **Configure the application:**
    - Rename `config.example.toml` to `config.toml` (or create your own `config.toml`).
    - Update the `config.toml` file with your database connection details, JWT secret key, and server port.

    ```toml
    # config.toml
    [Server]
    Port = "8848"

    [Database]
    Type = "sqlite3"
    DSN = "todo.db"

    [Jwt]
    SecretKey = "your-secret-key"
    Expire = 3600 # Expiration time in seconds (e.g., 1 hour)
    ```

### Running the Application

To run the API server, execute the following command from the root directory:

```sh
go run cmd/v1/main.go
```

The server will start on the port specified in your `config.toml` (e.g., `localhost:8848`).

## API Endpoints

The base path for all API endpoints is `/api/v1`.

### Swagger Documentation

For interactive API documentation, navigate to:
`http://localhost:8848/swagger/index.html`

### User Management

-   **Register a new user:**
    -   `POST /user/register`
    -   **Body**: `{"name": "YourName", "account": "user_account", "password": "your_password"}`

-   **Login:**
    -   `POST /user/login`
    -   **Body**: `{"account": "user_account", "password": "your_password"}`
    -   **Returns**: A JWT token for authenticating subsequent requests.

-   **Update User Info:**
    -   `PUT /user/update`
    -   **Authorization**: `Bearer <token>`
    -   **Body**: `{"name": "NewName", "password": "new_password"}`

### To-Do Management

*All To-Do endpoints require `Authorization: Bearer <token>` header.*

-   **Create a new to-do item:**
    -   `POST /todos/add`
    -   **Body**: `{"title": "My Task", "description": "Task details", "status": 0, "end_time": "2025-12-31T23:59:59Z"}`

-   **List all of the user's to-do items:**
    -   `GET /todos/list`

-   **Get a specific to-do item by ID:**
    -   `GET /todos/get/:id`

-   **Update a to-do item:**
    -   `PUT /todos/update`
    -   **Body**: `{"title": "Updated Title", "description": "Updated desc", "status": 1}`

-   **Delete a to-do item by ID:**
    -   `DELETE /todos/delete/:id`

## Project Structure

The project follows a standard Go application layout:

-   `cmd/`: Main application entry points.
    -   `v1/main.go`: The main function to start the API server.
-   `internal/`: Private application code, not importable by other projects.
    -   `database/`: Database initialization.
    -   `dto/`: Data Transfer Objects for request bodies.
    -   `handler/`: Gin handlers to process HTTP requests.
    -   `middleware/`: Gin middleware (e.g., JWT authentication).
    -   `model/`: Database models/entities.
    -   `repository/`: Data access layer for database operations.
    -   `router/`: API route definitions.
    -   `service/`: Business logic layer.
    -   `vo/`: View Objects for API responses.
-   `pkg/`: Public library code, can be imported by external projects.
    -   `jwtUtil/`: JWT generation and parsing utilities.
    -   `logUtil/`: Logger configuration.
-   `docs/`: Swagger documentation files.
-   `config.toml`: Application configuration file.
-   `go.mod`: Go module definitions and dependencies.

