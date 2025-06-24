package workshift

import (
	retailoutlet_models "my-market-server/main/retail_outlet/models"
	"time"
)

type WorkShiftModel struct {
	ID             int `gorm:"primaryKey; autoIncrement"`
	CreatedAt      time.Time
	Cash           float32
	Cashless       float32
	CashRegister   float32                               `gorm:"column:cash_register"`
	RetailOutletId int                                   `gorm:"column:retail_outlet_id"`
	RetailOutlet   retailoutlet_models.RetailOutletModel `gorm:"foreignKey:RetailOutletId"`
	CardTransfers  []CardTransferModel                   `gorm:"foreignKey:WorkShiftId"`
	Expenses       []ExpenseModel                        `gorm:"foreignKey:WorkShiftId"`
}
