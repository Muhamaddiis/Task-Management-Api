package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Muhamaddiis/Task-Management-Api/database"
	"github.com/Muhamaddiis/Task-Management-Api/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var taskReq models.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&taskReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate due date is in the future
	if taskReq.DueDate.Before(time.Now()) {
		http.Error(w, "Due date must be in the future", http.StatusBadRequest)
		return
	}

	task := models.Task{
		Title:       taskReq.Title,
		Description: taskReq.Description,
		Status:      taskReq.Status,
		DueDate:     taskReq.DueDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if task.Status == "" {
		task.Status = "pending"
	}

	// Check if title already exists
	var existingTask models.Task
	if err := database.DB.Where("title = ?", task.Title).First(&existingTask).Error; err == nil {
		http.Error(w, "Task with this title already exists", http.StatusConflict)
		return
	}

	if err := database.DB.Create(&task).Error; err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	response := models.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	queryParams := models.TaskQueryParams{
		Status:  r.URL.Query().Get("status"),
		DueDate: r.URL.Query().Get("due_date"),
		Search:  r.URL.Query().Get("search"),
		Page:    1,
		Limit:   10,
	}

	// Parse pagination parameters
	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			queryParams.Page = page
		}
	}

	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 && limit <= 100 {
			queryParams.Limit = limit
		}
	}

	db := database.DB.Model(&models.Task{})

	// Apply filters
	if queryParams.Status != "" {
		db = db.Where("status = ?", queryParams.Status)
	}

	if queryParams.DueDate != "" {
		if dueDate, err := time.Parse("2006-01-02", queryParams.DueDate); err == nil {
			db = db.Where("DATE(due_date) = ?", dueDate.Format("2006-01-02"))
		}
	}

	if queryParams.Search != "" {
		db = db.Where("title ILIKE ?", "%"+queryParams.Search+"%")
	}

	// Calculate offset
	offset := (queryParams.Page - 1) * queryParams.Limit

	var total int64
	db.Count(&total)

	var tasks []models.Task
	if err := db.Offset(offset).Limit(queryParams.Limit).Find(&tasks).Error; err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	// Prepare response
	var response []models.TaskResponse
	for _, task := range tasks {
		response = append(response, models.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			DueDate:     task.DueDate,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}

	// Add pagination metadata
	pagination := map[string]interface{}{
		"page":       queryParams.Page,
		"limit":      queryParams.Limit,
		"total":      total,
		"totalPages": (int(total) + queryParams.Limit - 1) / queryParams.Limit,
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"tasks":      response,
		"pagination": pagination,
	})
}
