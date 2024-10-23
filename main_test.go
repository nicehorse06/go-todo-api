package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestCreateTask tests the POST /tasks API
func TestCreateTask(t *testing.T) {
	router := setupRouter()

	task := Task{
		Title:       "Test Task",
		Description: "This is a test task",
		DueDate:     parseTime("2024-12-31T23:59:59Z"),
	}

	w := httptest.NewRecorder()
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdTask Task
	json.Unmarshal(w.Body.Bytes(), &createdTask)
	assert.Equal(t, "Test Task", createdTask.Title)
}

// TestGetAllTasks tests the GET /tasks API
func TestGetAllTasks(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var tasks []Task
	json.Unmarshal(w.Body.Bytes(), &tasks)
	assert.GreaterOrEqual(t, len(tasks), 0)
}

// TestGetTaskByID tests the GET /tasks/:id API
func TestGetTaskByID(t *testing.T) {
	router := setupRouter()

	// First, create a task to get its ID
	task := Task{
		Title:       "Test Task for Get",
		Description: "This is a task for testing Get",
		DueDate:     parseTime("2024-12-31T23:59:59Z"),
	}

	w := httptest.NewRecorder()
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var createdTask Task
	json.Unmarshal(w.Body.Bytes(), &createdTask)

	// Now, get the task by ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/tasks/"+strconv.Itoa(createdTask.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetchedTask Task
	json.Unmarshal(w.Body.Bytes(), &fetchedTask)
	assert.Equal(t, createdTask.ID, fetchedTask.ID)
}

// TestUpdateTask tests the PUT /tasks/:id API
func TestUpdateTask(t *testing.T) {
	router := setupRouter()

	// First, create a task to update
	task := Task{
		Title:       "Test Task for Update",
		Description: "This is a task for testing Update",
		DueDate:     parseTime("2024-12-31T23:59:59Z"),
	}

	w := httptest.NewRecorder()
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var createdTask Task
	json.Unmarshal(w.Body.Bytes(), &createdTask)

	// Now, update the task
	updatedTask := Task{
		Title:       "Updated Task Title",
		Description: "Updated description",
		DueDate:     parseTime("2024-12-31T23:59:59Z"),
		Status:      "in progress",
	}

	w = httptest.NewRecorder()
	body, _ = json.Marshal(updatedTask)
	req, _ = http.NewRequest("PUT", "/tasks/"+strconv.Itoa(createdTask.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var taskAfterUpdate Task
	json.Unmarshal(w.Body.Bytes(), &taskAfterUpdate)
	assert.Equal(t, "Updated Task Title", taskAfterUpdate.Title)
	assert.Equal(t, "in progress", taskAfterUpdate.Status)
}

// TestDeleteTask tests the DELETE /tasks/:id API
func TestDeleteTask(t *testing.T) {
	router := setupRouter()

	// First, create a task to delete
	task := Task{
		Title:       "Test Task for Delete",
		Description: "This is a task for testing Delete",
		DueDate:     parseTime("2024-12-31T23:59:59Z"),
	}

	w := httptest.NewRecorder()
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var createdTask Task
	json.Unmarshal(w.Body.Bytes(), &createdTask)

	// Now, delete the task by ID
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(createdTask.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Try to get the task again to verify deletion
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/tasks/"+strconv.Itoa(createdTask.ID), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestMarkTaskComplete tests the PATCH /tasks/:id/complete API
func TestMarkTaskComplete(t *testing.T) {
	router := setupRouter()

	// First, create a task to mark complete
	task := Task{
		Title:       "Test Task for Complete",
		Description: "This is a task for testing Complete",
		DueDate:     parseTime("2024-12-31T23:59:59Z"),
	}

	w := httptest.NewRecorder()
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var createdTask Task
	json.Unmarshal(w.Body.Bytes(), &createdTask)

	// Now, mark the task as complete
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PATCH", "/tasks/"+strconv.Itoa(createdTask.ID)+"/complete", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var completedTask Task
	json.Unmarshal(w.Body.Bytes(), &completedTask)
	assert.Equal(t, "complete", completedTask.Status)
}

// Helper function to set up the router
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/tasks", createTask)
	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getTaskByID)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.PATCH("/tasks/:id/complete", markTaskComplete)

	return router
}

// Helper function to parse time
func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}
