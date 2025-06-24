package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"task/pkg/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	authHeader = "Authorization"
	authToken  = "Bearer secret-token"
)

func TestGetEmptyTasks(t *testing.T) {
	mux := router.SetupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/tasks", nil)
	req.Header.Set(authHeader, authToken)
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]\n", w.Body.String())
}

func TestCreateTask(t *testing.T) {
	mux := router.SetupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"name": "Task1", "status": 0}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(authHeader, authToken)
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp["id"])
}

func TestCreateTaskInvalidStatus(t *testing.T) {
	mux := router.SetupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"name": "Task1", "status": 2}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(authHeader, authToken)
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetTasks(t *testing.T) {
	mux := router.SetupRouter()
	wCreate := httptest.NewRecorder()
	createReq := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"name": "Task1", "status": 0}`))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set(authHeader, authToken)
	mux.ServeHTTP(wCreate, createReq)
	assert.Equal(t, http.StatusCreated, wCreate.Code)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/tasks", nil)
	req.Header.Set(authHeader, authToken)
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Task1")
}

func TestUpdateTask(t *testing.T) {
	mux := router.SetupRouter()
	wCreate := httptest.NewRecorder()
	createReq := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"name": "Task1", "status": 0}`))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set(authHeader, authToken)
	mux.ServeHTTP(wCreate, createReq)
	assert.Equal(t, http.StatusCreated, wCreate.Code)

	var created map[string]interface{}
	err := json.Unmarshal(wCreate.Body.Bytes(), &created)
	assert.NoError(t, err)
	taskID := created["id"].(string)
	assert.NotEmpty(t, taskID)

	wUpdate := httptest.NewRecorder()
	updateReq := httptest.NewRequest("PUT", "/tasks/"+taskID, strings.NewReader(`{"name": "Updated Task", "status": 1}`))

	updateReq.Header.Set("Content-Type", "application/json")
	updateReq.Header.Set(authHeader, authToken)
	mux.ServeHTTP(wUpdate, updateReq)
	assert.Equal(t, http.StatusOK, wUpdate.Code)
	assert.Contains(t, wUpdate.Body.String(), "Updated Task")
	assert.Contains(t, wUpdate.Body.String(), "1")
}

func TestUpdateTaskInvalidStatus(t *testing.T) {
	mux := router.SetupRouter()
	wCreate := httptest.NewRecorder()
	createReq := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"name": "Task1", "status": 0}`))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set(authHeader, authToken)
	mux.ServeHTTP(wCreate, createReq)
	assert.Equal(t, http.StatusCreated, wCreate.Code)

	var created map[string]interface{}
	err := json.Unmarshal(wCreate.Body.Bytes(), &created)
	assert.NoError(t, err)
	taskID := created["id"].(string)
	assert.NotEmpty(t, taskID)

	wUpdate := httptest.NewRecorder()
	updateReq := httptest.NewRequest("PUT", "/tasks/"+taskID, strings.NewReader(`{"name": "Updated Task", "status": 2}`))

	updateReq.Header.Set("Content-Type", "application/json")
	updateReq.Header.Set(authHeader, authToken)
	mux.ServeHTTP(wUpdate, updateReq)
	assert.Equal(t, http.StatusBadRequest, wUpdate.Code)
}

func TestDeleteNoExistTask(t *testing.T) {
	mux := router.SetupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/tasks/1", nil)
	req.Header.Set(authHeader, authToken)
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteTask(t *testing.T) {
	mux := router.SetupRouter()
	wCreate := httptest.NewRecorder()
	createReq := httptest.NewRequest("POST", "/tasks", strings.NewReader(`{"name": "Task1", "status": 0}`))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set(authHeader, authToken)
	mux.ServeHTTP(wCreate, createReq)
	assert.Equal(t, http.StatusCreated, wCreate.Code)

	var created map[string]interface{}
	err := json.Unmarshal(wCreate.Body.Bytes(), &created)
	assert.NoError(t, err)
	taskID := created["id"].(string)
	assert.NotEmpty(t, taskID)

	wDelete := httptest.NewRecorder()
	deleteReq := httptest.NewRequest("DELETE", "/tasks/"+taskID, nil)
	deleteReq.Header.Set(authHeader, authToken)
	mux.ServeHTTP(wDelete, deleteReq)
	assert.Equal(t, http.StatusNoContent, wDelete.Code)

	wNotFound := httptest.NewRecorder()
	repeatDeleteReq := httptest.NewRequest("DELETE", "/tasks/"+taskID, nil)
	repeatDeleteReq.Header.Set(authHeader, authToken)
	mux.ServeHTTP(wNotFound, repeatDeleteReq)
	assert.Equal(t, http.StatusNotFound, wNotFound.Code)
}

func TestUnauthorizedAccess(t *testing.T) {
	mux := router.SetupRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/tasks", nil)
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
