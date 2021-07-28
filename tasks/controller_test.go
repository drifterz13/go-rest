package tasks_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/drifterz13/go-rest-api/tasks"
	"github.com/gin-gonic/gin"
)

type MockTaskRepository struct{}

type MockTask struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type MockTaskResponse struct {
	Task MockTask `json:"task"`
}

type MockTasksResponse struct {
	Tasks []MockTask `json:"tasks"`
}

var (
	mockTask1 = tasks.Task{Title: "task1", Completed: false}
	mockTask2 = tasks.Task{Title: "task2", Completed: false}
	mockTasks = []tasks.Task{
		mockTask1,
		mockTask2,
	}
)

func (repo *MockTaskRepository) FindAll() (*[]tasks.Task, error) {
	return &mockTasks, nil
}

func (repo *MockTaskRepository) FindById(id string) (*tasks.Task, error) {
	return &mockTask1, nil
}

func (repo *MockTaskRepository) Create(task *tasks.Task) error {
	return nil
}

func (repo *MockTaskRepository) Last() (*tasks.Task, error) {
	return &mockTask1, nil
}

func (repo *MockTaskRepository) DeleteById(id string) error {
	return nil
}

func (repo *MockTaskRepository) UpdateById(id string, doc *tasks.UpdateTaskPayload) (*tasks.Task, error) {
	return &mockTask1, nil
}

func TestGetTasksHandler(t *testing.T) {
	t.Run("get tasks", func(t *testing.T) {
		router := gin.Default()
		repository := &MockTaskRepository{}
		controller := tasks.NewTaskController(repository)
		router.GET("/tasks", controller.GetTasks)

		req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		wantCode := http.StatusOK

		if w.Code != wantCode {
			t.Fatalf("status code, want: %d, got: %d", wantCode, w.Code)
		}
	})
}

func TestCreateTaskHandler(t *testing.T) {
	tt := []struct {
		name     string
		method   string
		wantCode int
		body     string
	}{
		{
			name:     "with valid payload",
			method:   http.MethodPost,
			wantCode: http.StatusOK,
			body:     `{"title":"Sleep","completed":false}`,
		},
		{
			name:     "with empty payload",
			method:   http.MethodPost,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "with invalid payload",
			method:   http.MethodPost,
			wantCode: http.StatusBadRequest,
			body:     `{"name":"Sleep","is_completed":false}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			router := gin.Default()
			repository := &MockTaskRepository{}
			controller := tasks.NewTaskController(repository)
			router.POST("/task", controller.CreateTask)

			req, _ := http.NewRequest(tc.method, "/task", strings.NewReader(tc.body))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if tc.wantCode != w.Code {
				t.Fatalf("status code, want: %d, got: %d", tc.wantCode, w.Code)
			}
		})
	}
}

func TestUpdateTaskHandler(t *testing.T) {
	tt := []struct {
		name     string
		method   string
		wantCode int
		body     string
	}{
		{
			name:     "with valid payload",
			method:   http.MethodPatch,
			wantCode: http.StatusOK,
			body:     `{"title":"Learn GO","completed":false}`,
		},
		{
			name:     "with partial payload",
			method:   http.MethodPatch,
			wantCode: http.StatusOK,
			body:     `{"completed":true}`,
		},
		{
			name:     "with empty payload",
			method:   http.MethodPatch,
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "with invalid payload",
			method:   http.MethodPatch,
			wantCode: 1,
			body:     `{"is_valid":false}`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			router := gin.Default()
			repository := &MockTaskRepository{}
			controller := tasks.NewTaskController(repository)
			router.PATCH("/task/:id", controller.UpdateTask)

			req, _ := http.NewRequest(tc.method, "/task/1", strings.NewReader(tc.body))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if tc.wantCode != w.Code {
				t.Fatalf("status code, want: %d, got: %d", tc.wantCode, w.Code)
			}
		})
	}
}

func TestDeleteTaskHandler(t *testing.T) {
	t.Run("delete task", func(t *testing.T) {
		router := gin.Default()
		repository := &MockTaskRepository{}
		controller := tasks.NewTaskController(repository)

		router.DELETE("/task/:id", controller.DeleteTask)

		req, _ := http.NewRequest(http.MethodDelete, "/task/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		wantCode := http.StatusNoContent

		if wantCode != w.Code {
			t.Fatalf("status code, want: %d, got: %d", wantCode, w.Code)
		}
	})
}
