package retailoutlet

import (
	"gorm.io/gorm"
)

type RetailOutletRepo struct {
	Repo *gorm.DB
}

type RetailOutletRepoService interface {
	FindAll() ([]RetailOutletModel, error)
	FindOne(id int) (RetailOutletModel, error)
	Create(dto RetailOutletModel) (RetailOutletModel, error)
}

func (r *RetailOutletRepo) Create(dto RetailOutletModel) (RetailOutletModel, error) {
	model := dto
	res := r.Repo.Create(&model)
	if res.Error != nil {
		return model, res.Error
	}
	model, err := r.FindOne(int(model.ID))
	return model, err
}

func (r *RetailOutletRepo) FindOne(id int) (RetailOutletModel, error) {
	var model RetailOutletModel
	res := r.Repo.Where("id = ?", id).Preload("Seo").First(&model)
	return model, res.Error
}

func (r *RetailOutletRepo) FindAll() ([]RetailOutletModel, error) {
	var model []RetailOutletModel
	res := r.Repo.Preload("Seo").Find(&model)
	return model, res.Error
}
