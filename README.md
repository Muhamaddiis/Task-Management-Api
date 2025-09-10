# Task Management System API

A simple RESTful API for managing tasks built with Golang, Gorilla Mux, and GORM.

## Features

- CRUD operations for tasks
- Task validation (required fields, unique title, future due date)
- Filtering by status and due date
- Search by title
- Pagination
- PostgreSQL database

## Setup Instructions

### Prerequisites

- Go 1.21 or higher
- PostgreSQL

### Installation

1. Clone the repository:
    ```bash
    git clone <repository-url>
    cd task-management-api

2. Install dependencies:
    ```bash
    go mod download
3. Set environment variables
    ```bash
    export DATABASE_URL="host=localhost user=postgres password=yourpassword dbname=task_management port=5432 sslmode=disable"
4. Set up PostgreSQL database:
    ```bash
    createdb taskdb
    # or via psql:
    # psql -U postgres -c "CREATE DATABASE taskdb;"
5. Migration
     The application will automatically run migrations on startup
6. Start the server
    ```bash
    go run main.go

## API Endpoints
### Create a Task
* POST /api/v1/tasks

* Body:
    ```json
    {
        "title": "Task title",
        "description": "Task description",
        "status": "pending",
        "due_date": "2023-12-31T23:59:59Z"
    }
### Get All Tasks
* GET /api/v1/tasks

* Query parameters
    * status - Filter by status (pending, in_progress, completed)

    * due_date - Filter by due date (YYYY-MM-DD format)

    * search - Search by title

    * page - Page number (default: 1)

    * limit - Items per page (default: 10, max: 100)

### Get a specific Task
* GET /api/v1/tasks/{id}

### Update a Task
* PUT /api/v1/tasks/{id}

* Body same as createTask

### Delete a Task
* DELETE /api/v1/tasks/{id}

### Validation rules
* Title is required and must be unique

* Status must be one of: pending, in_progress, completed

* Due date must be in the future

### Example Usage
1. Create a Task:

    ```bash
        curl -X POST http://localhost:8080/api/v1/tasks \
        -H "Content-Type: application/json" \
        -d '{
                "title": "Complete API project",
                "description": "Finish the task management API",
                "status": "pending",
                "due_date": "2023-12-31T23:59:59Z"
            }' 
2. Get all tasks with pagination
    ```bash
    curl "http://localhost:8080/api/v1/tasks?page=1&limit=10"

3. Filter task by status
    ```bash
    curl "http://localhost:8080/api/v1/tasks?status=pending"
    
4. Search task by Title
    ```bash
    curl "http://localhost:8080/api/v1/tasks?search=complete"






This implementation provides a complete Task Management System API with all the requested features:
- CRUD operations for tasks
- Validation for required fields and data types
- PostgreSQL database with migrations
- Filtering by status and due date
- Pagination
- Search functionality
- Proper error handling

The API follows RESTful principles and includes proper validation, error handling, and documentation.