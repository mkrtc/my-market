package workshift

type ExpenseModel struct {
	ID          int             `gorm:"primaryKey; autoIncrement" json:"id"`
	Article     string          `json:"article"`
	Debit       float32         `json:"debit"`
	Credit      float32         `json:"credit"`
	Payed       bool            `json:"payed"`
	WorkShiftId int             `gorm:"column:work_shit_id" json:"work_shift_id"`
	WorkShift   *WorkShiftModel `gorm:"foreignKey:WorkShiftId" json:"work_shift"`
}
