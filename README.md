# GOLANG TO-DO LIST

A simple todo list application with user authentication.

## instalation

1. Clone the repository:
   ```
   git clone https://github.com/radeqq007/golang-todo-list/
   cd golang-todo-list
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Set up the database according to `internal/database/schema.sql`.
4. Run the main.go file:
   ```
   go run ./cmd/app/main.go
   ```

## Configuration

This app uses mariaDB v8.0.30.
