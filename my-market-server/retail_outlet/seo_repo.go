package retailoutlet

import (
	"gorm.io/gorm"
)

type SeoRepo struct {
	Repo *gorm.DB
}

type SeoRepoService interface {
	FindById(id int, model *SeoModel) error
	FindAll(model *[]SeoModel) error
	Create(model *SeoModel) error
}

func (r *SeoRepo) FindById(id int, model *SeoModel) error {
	res := r.Repo.Where("id = ?", id).First(&model)
	return res.Error
}

func (r *SeoRepo) FindAll(model *[]SeoModel) error {
	res := r.Repo.Preload("RetailOutlets").Find(&model)
	return res.Error
}

func (r *SeoRepo) Create(model *SeoModel) error {
	res := r.Repo.Create(&model)
	return res.Error
}
