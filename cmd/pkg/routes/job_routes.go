package routes

import (
	"github.com/gorilla/mux"
	"github.com/moheddine-belhaj/Job_app/pkg/controllers"
)

var RegisterJobRoutes = func(router *mux.Router) {
	router.HandleFunc("/job/create", controllers.CreateJob).Methods("POST")
	router.HandleFunc("/job/all", controllers.GetJobs).Methods("GET")
	router.HandleFunc("/job/{jobId}", controllers.GetJobById).Methods("GET")
	router.HandleFunc("/job/update/{jobId}", controllers.UpdateJob).Methods("PUT")
	router.HandleFunc("/job/delete/{jobId}", controllers.DeleteJob).Methods("DELETE")
}
