package retailoutlet

import (
	"gorm.io/gorm"
)

type RetailOutletRepo struct {
	Repo *gorm.DB
}

type RetailOutletRepoService interface {
	Create(model *RetailOutletModel) error
	FindOne(id int, retailOutlet *RetailOutletModel) error
	FindAll(retailOutlets *[]RetailOutletModel) error
}

func (r *RetailOutletRepo) Create(model *RetailOutletModel) error {
	result := r.Repo.Create(&model)
	return result.Error
}

func (r *RetailOutletRepo) FindOne(id int, retailOutlet *RetailOutletModel) error {
	result := r.Repo.Where("id = ?", id).First(&retailOutlet)
	return result.Error
}

func (r *RetailOutletRepo) FindAll(retailOutlets *[]RetailOutletModel) error {
	result := r.Repo.Find(&retailOutlets)
	return result.Error
}
