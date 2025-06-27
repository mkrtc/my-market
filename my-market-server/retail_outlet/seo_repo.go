package retailoutlet

import (
	"gorm.io/gorm"
)

type SeoRepo struct {
	Repo *gorm.DB
}

type SeoRepoService interface {
	FindById(id int) (SeoModel, error)
	FindAll() ([]SeoModel, error)
	Create(dto SeoModel) (SeoModel, error)
}

func (r *SeoRepo) FindById(id int) (SeoModel, error) {
	var model SeoModel
	res := r.Repo.Where("id = ?", id).First(&model)
	return model, res.Error
}

func (r *SeoRepo) FindAll() ([]SeoModel, error) {
	var model []SeoModel
	res := r.Repo.Preload("RetailOutlets").Find(&model)
	return model, res.Error
}

func (r *SeoRepo) Create(dto SeoModel) (SeoModel, error) {
	model := dto
	res := r.Repo.Create(&model)
	return model, res.Error
}
