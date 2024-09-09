package repository

import (
	"github.com/user_app/model"
	"gorm.io/gorm"
)

type JobRepository struct {
	JobRepo *Reposiotry
}

func NewJobRepository(jobrepo *Reposiotry) *JobRepository {
	return &JobRepository{
		JobRepo: jobrepo,
	}
}

func (j *JobRepository) CreateJob(job *model.Job) error {

	err := j.JobRepo.DB.Create(&job).Error

	if err != nil {
		return err
	}

	return nil

}

// for student
func (j *JobRepository) GetAllJobs(keyword string) ([]model.Job, error) {
	var jobs []model.Job

	err := j.JobRepo.DB.Preload("Company").Where("title ILIKE ? OR description ILIKE ?", "%"+keyword+"%", "%"+keyword+"%").Order("created_at DESC").Find(&jobs).Error

	if err != nil {
		return nil, err
	}

	return jobs, nil

}

// for student
func (j *JobRepository) GetJobByID(id uint) (model.Job, error) {
	var job model.Job

	// err := j.JobRepo.DB.Model(model.Job{}).Preload("Applications").Where("id = ?", id).Find(&job).Error
	err := j.JobRepo.DB.Model(model.Job{}).
		Preload("Applications", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, job_id, user_id, status")
		}).Where("id = ?", id).Find(&job).Error

	if err != nil {
		return model.Job{}, err
	}

	return job, nil

}

func (j *JobRepository) UpdateJob(job model.Job, id uint) error {

	err := j.JobRepo.DB.Model(model.Job{}).Where("id = ?", id).Updates(job).Error

	if err != nil {
		return err
	}

	return nil

}

func (j *JobRepository) DeleteJob(job model.Job) error {

	err := j.JobRepo.DB.Delete(&job).Error

	if err != nil {
		return err
	}

	return nil

}

func (j *JobRepository) GetJobByName(name string) (model.Job, error) {

	var job model.Job

	err := j.JobRepo.DB.Model(model.Job{}).Where("title = ?", name).Find(&job).Error

	if err != nil {
		return model.Job{}, err
	}

	return job, nil

}

func (j *JobRepository) GetJobsByCompanyID(id uint) ([]model.Job, error) {

	var jobs []model.Job

	err := j.JobRepo.DB.Model(model.Job{}).Where("company_id = ?", id).Find(&jobs).Error

	if err != nil {
		return nil, err
	}

	return jobs, nil

}
func (j *JobRepository) GetAdminJobs(uid uint) ([]model.Job, error) {

	var jobs []model.Job

	err := j.JobRepo.DB.Model(model.Job{}).Preload("Company").Where("created_by_id = ?", uid).Order("created_at DESC").Find(&jobs).Error

	if err != nil {
		return nil, err
	}

	return jobs, nil

}
