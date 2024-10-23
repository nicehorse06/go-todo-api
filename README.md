
# Todo List API

This project is a simple RESTful API for managing a Todo List built using Golang, Gin, and a global in-memory store (which can be later replaced with Postgres or Redis).

## Installation

If you are new to the project, follow these steps to set up the environment and install the necessary dependencies.

### Prerequisites
- **Go 1.16 or higher** is required. You can install Go from [the official website](https://golang.org/dl/).

### Setup Instructions

1. **Clone the repository**
   Clone the project from GitHub:
   ```bash
   git clone git@github.com:nicehorse06/todo-list-api.git
   cd todo-list-api
   ```

2. **Initialize the Go module**
   Ensure the project is using Go modules. Run the following command to initialize Go modules if it hasn't been done yet:
   ```bash
   go mod init github.com/nicehorse06/todo-list-api
   ```

3. **Install dependencies**
   Use `go get` to install all required dependencies for this project, including the Gin framework and testing libraries:
   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u github.com/stretchr/testify/assert
   ```

   Alternatively, if the `go.mod` file is present, run the following to install dependencies:
   ```bash
   go mod tidy
   ```

4. **Run the application**
   After installing dependencies, you can run the project with the following command:
   ```bash
   go run main.go
   ```

5. **Run the tests**
   To run the test cases and ensure everything is working properly, execute:
   ```bash
   go test -v
   ```

## go CLI Commands

Here are some useful Go CLI commands to work with the project.

### Run the API server

Use the following command to start the API server:
```bash
go run main.go
```

This will start the server at `localhost:8080`, and you can access the endpoints such as `http://localhost:8080/tasks`.

### Run the tests

Use the following command to run all the test cases:
```bash
go test -v
```

The `-v` flag provides verbose output, showing detailed information for each test case.

## Endpoints

The following endpoints are available:

1. **Create a Task**
   - **POST** `/tasks`
   - Body parameters: `title`, `description`, `due_date`
   - Response: Returns the created task details

2. **Get All Tasks**
   - **GET** `/tasks`
   - Response: Returns a list of all tasks

3. **Get a Single Task**
   - **GET** `/tasks/:id`
   - Parameters: `id` (the task ID)
   - Response: Returns the task details by its ID

4. **Update a Task**
   - **PUT** `/tasks/:id`
   - Parameters: `id` (the task ID)
   - Body parameters: `title`, `description`, `status`
   - Response: Returns the updated task details

5. **Delete a Task**
   - **DELETE** `/tasks/:id`
   - Parameters: `id` (the task ID)
   - Response: Deletes the task and returns a success message

6. **Mark Task as Complete**
   - **PATCH** `/tasks/:id/complete`
   - Parameters: `id` (the task ID)
   - Response: Marks the task as complete and returns the updated task details
