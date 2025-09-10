#!/bin/bash

API_URL="http://localhost:8080/api/v1/tasks"

# Array of task data
tasks=(
  '{"title":"Review Code Quality","description":"Perform code review and refactor where necessary","status":"pending","due_date":"2025-11-20T18:00:00Z"}'
  '{"title":"Setup CI/CD Pipeline","description":"Configure continuous integration and deployment pipeline","status":"completed","due_date":"2025-10-31T17:30:00Z"}'
  '{"title":"Write Unit Tests","description":"Create comprehensive test suite for all API endpoints","status":"in_progress","due_date":"2025-09-25T16:45:00Z"}'
  '{"title":"Database Optimization","description":"Optimize database queries and add indexes","status":"pending","due_date":"2025-09-01T14:00:00Z"}'
  '{"title":"User Authentication","description":"Implement JWT-based user authentication system","status":"pending","due_date":"2025-09-28T12:00:00Z"}'
  '{"title":"Frontend Integration","description":"Connect React frontend to the API endpoints","status":"in_progress","due_date":"2025-12-22T15:30:00Z"}'
  '{"title":"Deploy to Production","description":"Deploy the application to production environment","status":"pending","due_date":"2025-11-05T10:00:00Z"}'
  '{"title":"Performance Testing","description":"Run load and stress tests on the API","status":"completed","due_date":"2025-9-25T11:45:00Z"}'
  '{"title":"Security Audit","description":"Conduct security review and implement fixes","status":"in_progress","due_date":"2025-10-18T09:00:00Z"}'
)

echo "Populating database with sample tasks..."

for task in "${tasks[@]}"; do
  echo "Adding task: $task"
  curl -X POST "$API_URL" \
    -H "Content-Type: application/json" \
    -d "$task"
  echo -e "\n"
done

echo "Database population completed!"