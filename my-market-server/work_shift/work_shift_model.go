package workshift

import (
	rt "my-market-server/main/retail_outlet"
	"time"
)

type WorkShiftModel struct {
	ID             int                   `gorm:"primaryKey; autoIncrement" json:"id"`
	CreatedAt      time.Time             `json:"createdAt"`
	Cash           float32               `json:"cash"`
	Cashless       float32               `json:"cash_less"`
	CashRegister   float32               `gorm:"column:cash_register" json:"cash_register"`
	RetailOutletId int                   `gorm:"column:retail_outlet_id" json:"retail_outlet_id"`
	RetailOutlet   *rt.RetailOutletModel `gorm:"foreignKey:RetailOutletId" json:"retail_outlet"`
	CardTransfers  []CardTransferModel   `gorm:"foreignKey:WorkShiftId" json:"card_transfers"`
	Expenses       []ExpenseModel        `gorm:"foreignKey:WorkShiftId" json:"expenses"`
}
