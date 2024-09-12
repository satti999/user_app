package repository

import (
	"errors"
	"fmt"

	"github.com/user_app/model"
	"gorm.io/gorm"
)

type ApplicationRepository struct {
	ApplicationRepo *Reposiotry
}

func NewApplicationRepository(applicationrepo *Reposiotry) *ApplicationRepository {
	return &ApplicationRepository{
		ApplicationRepo: applicationrepo,
	}
}

func (repo *ApplicationRepository) ApplyJob(application *model.Application) error {
	result := repo.CheckExistingApplication(application.UserID, application.JobID)
	var job model.Job
	if result {
		return errors.New("you have already applied for this job")
	}
	res := repo.ApplicationRepo.DB.First(&job, application.JobID)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		// Return a custom error message if the job is not found
		return fmt.Errorf("job with ID %d not found", application.JobID)
	}

	err := repo.ApplicationRepo.DB.Model(model.Application{}).Create(&application).Error

	if err != nil {

		return err
	}

	return nil

}

func (repo *ApplicationRepository) GetAppliedJobs(id uint) ([]model.Application, error) {

	var applications []model.Application

	err := repo.ApplicationRepo.DB.Model(model.Application{}).Preload("Job").Preload("Job.Company").Where("user_id = ? ", id).Find(&applications).Error
     
	if err != nil {

		return nil, err

	}

	return applications, nil

}

func (repo *ApplicationRepository) CheckExistingApplication(id uint, jobid uint) bool {

	var application model.Application

	err := repo.ApplicationRepo.DB.Model(model.Application{}).Where("user_id = ? AND job_id = ?", id, jobid).Find(&application).Error

	if err != nil {

		return false
	}
	if application.ID == 0 {

		return false
	}
	return true

}
func (repo *ApplicationRepository) UpdateStatus(status string, id uint) error {

	err := repo.ApplicationRepo.DB.Model(model.Application{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil

}

func (repo *ApplicationRepository) GetApplication(id uint) ([]model.Application, error) {
	applica := []model.Application{}

	err := repo.ApplicationRepo.DB.Model(model.Application{}).Preload("User").Preload("User.Profile").Where("job_id = ?", id).Find(&applica).Error

	if err != nil {
		return []model.Application{}, err

	}

	return applica, nil

}
