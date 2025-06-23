package retailoutlet

import (
	"gorm.io/gorm"
)

type RetailOutletRepo struct {
	Repo *gorm.DB
}

type RetailOutletRepoService interface {
	Create(dto *RetailOutletModel) error
	FindOne(id int, retailOutlet *RetailOutletModel) error
	FindAll(retailOutlets *[]RetailOutletModel) error
}

func (r *RetailOutletRepo) Create(dto *RetailOutletModel) error {
	result := r.Repo.Create(&dto)
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
