package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Task struct to define task fields
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// Global variable to act as a mock database
var tasks = make(map[int]Task)
var taskID = 1
var mu sync.Mutex // Mutex for thread-safe access to global variables

func main() {
	r := gin.Default()

	// Create a Task
	r.POST("/tasks", createTask)

	// Get all Tasks
	r.GET("/tasks", getAllTasks)

	// Get a single Task by ID
	r.GET("/tasks/:id", getTaskByID)

	// Update a Task by ID
	r.PUT("/tasks/:id", updateTask)

	// Delete a Task by ID
	r.DELETE("/tasks/:id", deleteTask)

	// Mark a Task as Complete
	r.PATCH("/tasks/:id/complete", markTaskComplete)

	// Start server on port 8080
	r.Run(":8080")
}

// Handler to create a new task
func createTask(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	newTask.ID = taskID
	newTask.Status = "pending"
	tasks[taskID] = newTask
	taskID++
	mu.Unlock()

	c.JSON(http.StatusCreated, newTask)
}

// Handler to get all tasks
func getAllTasks(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var taskList []Task
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	c.JSON(http.StatusOK, taskList)
}

// Handler to get a task by ID
func getTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	mu.Lock()
	task, exists := tasks[id]
	mu.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Handler to update a task by ID
func updateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	task, exists := tasks[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Update task fields
	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.DueDate = updatedTask.DueDate
	task.Status = updatedTask.Status

	tasks[id] = task
	c.JSON(http.StatusOK, task)
}

// Handler to delete a task by ID
func deleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	_, exists := tasks[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	delete(tasks, id)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

// Handler to mark a task as complete
func markTaskComplete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	task, exists := tasks[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	task.Status = "complete"
	tasks[id] = task

	c.JSON(http.StatusOK, task)
}
