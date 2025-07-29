# Simple User CRUD API

A simple Go-based CRUD (Create, Read, Update, Delete) application for managing users with in-memory storage.

## Features

- Create users with name and email
- Get all users
- Get a specific user by ID
- Update user information
- Delete users
- In-memory storage (data is lost when server restarts)
- Thread-safe operations
- RESTful API design

## Prerequisites

- Go 1.21 or later

## Installation

1. Clone or download this repository
2. Navigate to the project directory
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Running the Application

Start the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### 1. Get API Documentation
- **GET** `/`
- Returns a simple HTML page with API documentation

### 2. Create User
- **POST** `/users`
- **Body:**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com"
  }
  ```
- **Response:** Created user with ID and timestamps

### 3. Get All Users
- **GET** `/users`
- **Response:** Array of all users

### 4. Get User by ID
- **GET** `/users/{id}`
- **Response:** Specific user or 404 if not found

### 5. Update User
- **PUT** `/users/{id}`
- **Body:**
  ```json
  {
    "name": "John Updated",
    "email": "john.updated@example.com"
  }
  ```
- **Response:** Updated user or 404 if not found

### 6. Delete User
- **DELETE** `/users/{id}`
- **Response:** Success message or 404 if not found

## Example Usage

### Using curl

1. **Create a user:**
   ```bash
   curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice Johnson", "email": "alice@example.com"}'
   ```

2. **Get all users:**
   ```bash
   curl http://localhost:8080/users
   ```

3. **Get a specific user:**
   ```bash
   curl http://localhost:8080/users/1
   ```

4. **Update a user:**
   ```bash
   curl -X PUT http://localhost:8080/users/1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Alice Smith", "email": "alice.smith@example.com"}'
   ```

5. **Delete a user:**
   ```bash
   curl -X DELETE http://localhost:8080/users/1
   ```

## Sample Data

The application starts with two sample users:
- John Doe (john@example.com)
- Jane Smith (jane@example.com)

## Response Format

All API responses follow this format:
```json
{
  "success": true,
  "message": "Optional message",
  "data": {
    // Response data
  }
}
```

## Error Handling

- **400 Bad Request:** Invalid request body or parameters
- **404 Not Found:** User with specified ID doesn't exist
- **500 Internal Server Error:** Server-side errors

## Notes

- Data is stored in memory and will be lost when the server restarts
- User IDs are auto-incremented starting from 1
- All operations are thread-safe using read-write mutexes
- Timestamps are automatically managed (created_at, updated_at) 