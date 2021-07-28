package tasks_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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

func (repo *MockTaskRepository) UpdateById(id string, doc *tasks.UpdateTaskDoc) (*tasks.Task, error) {
	return &mockTask1, nil
}

func TestTaskController(t *testing.T) {
	router := gin.Default()
	controller := &tasks.TaskController{Repo: &MockTaskRepository{}}

	taskRoutes := tasks.NewTaskRoutes(router, controller)
	taskRoutes.Register()

	tt := []struct {
		name       string
		method     string
		url        string
		statusCode int
	}{
		{
			name:       "get task",
			url:        "/task/1",
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
		{
			name:       "get tasks",
			url:        "/tasks",
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf(fmt.Sprintf("case %s\n", tc.name))

			req, _ := http.NewRequest(tc.method, tc.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tc.statusCode {
				t.Errorf("Want status: %d but got: %d\n", tc.statusCode, w.Code)
			}

			var result interface{}
			resp, _ := io.ReadAll(w.Body)

			json.Unmarshal(resp, &result)

			t.Logf("status: %d", w.Code)
			t.Logf("read value: %v", string(resp[:]))
			// t.Logf("response: %v", result)
		})
	}

}
