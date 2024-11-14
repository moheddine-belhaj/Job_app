package models

import (
	"github.com/moheddine-belhaj/Job_app/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

// Job model matching the provided structure
type Job struct {
	gorm.Model
	Title       string  `json:"title"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Salary      string  `json:"salary"`
	Company     Company `json:"company" gorm:"embedded"`
}

// Company model to handle embedded fields within Job
type Company struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ContactEmail string `json:"contactEmail"`
	ContactPhone string `json:"contactPhone"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Job{})
}

func (j *Job) CreateJob() *Job {
	db.Create(&j)
	return j
}

func GetAllJobs() []Job {
	var jobs []Job
	db.Find(&jobs)
	return jobs
}

func GetJobById(Id int64) (*Job, *gorm.DB) {
	var getJob Job
	db := db.Where("ID=?", Id).Find(&getJob)
	return &getJob, db
}

func DeleteJob(id int64) Job {
	var job Job
	db.Delete(&job, id)
	return job
}

// func UpdateJob(b *Job) (*Job, *gorm.DB) {
// 	result := db.Save(b)
// 	return b, result
// }
