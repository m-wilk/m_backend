package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"userId"`
}

type Response struct {
	Todos []Todo `json:"todos"`
	Total int    `json:"total"`
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
}

var (
	todos     []Todo
	todoMutex sync.Mutex
	nextID    = 1
)

// Handler to get todos
func (h *Handler) getTodos(c echo.Context) error {
	// Pagination parameters
	skip, _ := strconv.Atoi(c.QueryParam("skip"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 30
	}

	todoMutex.Lock()
	defer todoMutex.Unlock()

	// Limit data
	start := skip
	if start > len(todos) {
		start = len(todos)
	}
	end := start + limit
	if end > len(todos) {
		end = len(todos)
	}

	// Prepare response
	response := Response{
		Todos: todos[start:end],
		Total: len(todos),
		Skip:  skip,
		Limit: limit,
	}

	return c.JSON(http.StatusOK, response)
}

// Handler to add a new todo
func (h *Handler) addTodo(c echo.Context) error {
	newTodo := new(Todo)
	if err := c.Bind(newTodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	// Ensure completed is either true or false
	if newTodo.Completed != true && newTodo.Completed != false {
		return echo.NewHTTPError(http.StatusBadRequest, "`completed` must be true or false")
	}

	todoMutex.Lock()
	defer todoMutex.Unlock()

	// Set ID and add to list
	newTodo.ID = nextID
	nextID++
	todos = append(todos, *newTodo)

	return c.JSON(http.StatusCreated, newTodo)
}

func (h *Handler) updateTodoCompleted(c echo.Context) error {
	// Parse ID from path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	// Parse new completed status from request body
	var updateData struct {
		Completed bool `json:"completed"`
	}
	if err := c.Bind(&updateData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	todoMutex.Lock()
	defer todoMutex.Unlock()

	// Find and update the todo
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = updateData.Completed
			return c.JSON(http.StatusOK, todos[i])
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Todo not found")
}

func (h *Handler) deleteTodo(c echo.Context) error {
	// Parse ID from path parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	todoMutex.Lock()
	defer todoMutex.Unlock()

	// Find and remove the todo
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...) // Remove the todo
			return c.JSON(http.StatusOK, map[string]string{"message": "Todo deleted"})
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Todo not found")
}

// Initialize with mock data
func init() {
	todos = []Todo{}

	nextID = len(todos) + 1
}
