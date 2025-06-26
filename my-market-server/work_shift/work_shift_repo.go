package workshift

import (
	"fmt"

	"gorm.io/gorm"
)

type WorkShiftRepo struct {
	Repo *gorm.DB
}

type WorkShiftRepoService interface {
	FindAll(model *[]WorkShiftModel, limit int) error
	FindById(id int, model *WorkShiftModel) error
	Create(model *WorkShiftModel) error
}

func (r *WorkShiftRepo) FindAll(model *[]WorkShiftModel, limit int, order string) error {
	res := r.Repo.Preload("RetailOutlet").Preload("Expenses").Preload("CardTransfers").Limit(limit).Order(fmt.Sprintf("created_at %s", order)).Find(&model)
	return res.Error
}

func (r *WorkShiftRepo) FindById(id int, model *WorkShiftModel) error {
	res := r.Repo.Where("id = ?", id).First(&model)
	return res.Error
}

func (r *WorkShiftRepo) Create(model *WorkShiftModel) error {
	res := r.Repo.Create(&model)
	return res.Error
}
