package workshift

import (
	"fmt"

	"gorm.io/gorm"
)

type WorkShiftRepo struct {
	Repo *gorm.DB
}

type WorkShiftRepoService interface {
	FindAll(limit int) ([]WorkShiftModel, error)
	FindById(id int) (WorkShiftModel, error)
	Create(dto WorkShiftModel) (WorkShiftModel, error)
}

func (r *WorkShiftRepo) FindAll(limit int, order string) ([]WorkShiftModel, error) {
	var model []WorkShiftModel
	res := r.Repo.Preload("RetailOutlet").Preload("Expenses").Preload("CardTransfers").Limit(limit).Order(fmt.Sprintf("created_at %s", order)).Find(&model)
	return model, res.Error
}

func (r *WorkShiftRepo) FindById(id int) (WorkShiftModel, error) {
	var model WorkShiftModel
	res := r.Repo.Where("id = ?", id).First(&model)
	return model, res.Error
}

func (r *WorkShiftRepo) Create(dto WorkShiftModel) (WorkShiftModel, error) {
	model := dto
	res := r.Repo.Create(&model)
	return model, res.Error
}
