package workshift

type ExpenseModel struct {
	ID          int `gorm:"primaryKey"`
	Article     string
	Debit       float32
	Credit      float32
	Payed       bool
	WorkShiftId int            `gorm:"column:work_shit_id"`
	WorkShift   WorkShiftModel `gorm:"foreignKey:WorkShiftId"`
}
