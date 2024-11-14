package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/moheddine-belhaj/Job_app/pkg/models"
	"github.com/moheddine-belhaj/Job_app/pkg/utils"
)

func GetJobs(w http.ResponseWriter, r *http.Request) {
    jobs := models.GetAllJobs()
    res, err := json.Marshal(jobs)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetJobById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    jobId := vars["jobId"]
    ID, err := strconv.ParseUint(jobId, 10, 32)
    if err != nil {
        http.Error(w, "Invalid job ID", http.StatusBadRequest)
        return
    }
    job, result := models.GetJobById(int64(ID))
    if result.Error != nil {
        http.Error(w, "Job not found", http.StatusNotFound)
        return
    }
    res, err := json.Marshal(job)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
    var createJob models.Job
    if err := utils.ParseBody(r, &createJob); err != nil {
        http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
        return
    }
    job := createJob.CreateJob()
    res, err := json.Marshal(job)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(res)
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    jobId := vars["jobId"]
    ID, err := strconv.ParseInt(jobId, 0, 0)
    if err != nil {
        http.Error(w, "Invalid job ID", http.StatusBadRequest)
        return
    }
    job := models.DeleteJob(ID)
    res, _ := json.Marshal(job)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func UpdateJob(w http.ResponseWriter, r *http.Request) {
    var updateJob = &models.Job{}
    utils.ParseBody(r, updateJob)
    vars := mux.Vars(r)
    jobId := vars["jobId"]
    ID, err := strconv.ParseInt(jobId, 0, 0)
    if err != nil {
        http.Error(w, "Invalid job ID", http.StatusBadRequest)
        return
    }
    jobDetails, db := models.GetJobById(ID)
    if updateJob.Title != "" {
        jobDetails.Title = updateJob.Title
    }
    if updateJob.Type != "" {
        jobDetails.Type = updateJob.Type
    }
    if updateJob.Description != "" {
        jobDetails.Description = updateJob.Description
    }
    if updateJob.Location != "" {
        jobDetails.Location = updateJob.Location
    }
    if updateJob.Salary != "" {
        jobDetails.Salary = updateJob.Salary
    }
    db.Save(&jobDetails)
    res, _ := json.Marshal(jobDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
