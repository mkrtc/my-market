package retailoutlet

import (
	retailoutlet_models "my-market-server/main/retail_outlet/models"

	"gorm.io/gorm"
)

type RetailOutletRepo struct {
	Repo *gorm.DB
}

type RetailOutletRepoService interface {
	Create(model *retailoutlet_models.RetailOutletModel) error
	FindOne(id int, retailOutlet *retailoutlet_models.RetailOutletModel) error
	FindAll(retailOutlets *[]retailoutlet_models.RetailOutletModel) error
}

func (r *RetailOutletRepo) Create(model *retailoutlet_models.RetailOutletModel) error {
	result := r.Repo.Create(&model)
	return result.Error
}

func (r *RetailOutletRepo) FindOne(id int, retailOutlet *retailoutlet_models.RetailOutletModel) error {
	result := r.Repo.Where("id = ?", id).First(&retailOutlet)
	return result.Error
}

func (r *RetailOutletRepo) FindAll(retailOutlets *[]retailoutlet_models.RetailOutletModel) error {
	result := r.Repo.Find(&retailOutlets)
	return result.Error
}
