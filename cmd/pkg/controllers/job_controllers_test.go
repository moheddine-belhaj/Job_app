package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/moheddine-belhaj/Job_app/pkg/models"
	// "github.com/moheddine-belhaj/Job_app/pkg/utils"
)

func TestGetJobs(t *testing.T) {
	req, err := http.NewRequest("GET", "/jobs", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetJobs)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}


func TestGetJobById(t *testing.T) {
	jobID := 1


	req, err := http.NewRequest("GET", "/jobs/"+strconv.Itoa(jobID), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/jobs/{jobId}", GetJobById)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

func TestCreateJob(t *testing.T) {
	job := models.Job{
		Title:       "Software Engineer",
		Type:        "Full-Time",
		Description: "A job description",
		Location:    "Remote",
		Salary:      "100000",
	}
	body, _ := json.Marshal(job)

	req, err := http.NewRequest("POST", "/jobs", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateJob)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}
func TestDeleteJob(t *testing.T) {
	jobID := 1
	req, err := http.NewRequest("DELETE", "/jobs/"+strconv.Itoa(jobID), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/jobs/{jobId}", DeleteJob)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

func TestUpdateJob(t *testing.T) {
	jobID := 1 
	updatedJob := models.Job{
		Title: "Updated Title",
	}
	body, _ := json.Marshal(updatedJob)

	req, err := http.NewRequest("PUT", "/jobs/"+strconv.Itoa(jobID), bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/jobs/{jobId}", UpdateJob)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}
