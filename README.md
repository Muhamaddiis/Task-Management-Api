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
    # The application will automatically run migrations on startup
6. Start the server
    ```bash
    go run main.go


This implementation provides a complete Task Management System API with all the requested features:
- CRUD operations for tasks
- Validation for required fields and data types
- PostgreSQL database with migrations
- Filtering by status and due date
- Pagination
- Search functionality
- Proper error handling

The API follows RESTful principles and includes proper validation, error handling, and documentation.