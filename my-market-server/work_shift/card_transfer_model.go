package workshift

type CardTransferModel struct {
	ID          int             `gorm:"primaryKey; autoIncrement" json:"id"`
	Sum         float32         `json:"sum"`
	WorkShiftId int             `gorm:"column:work_shift_id" json:"work_shift_id"`
	WorkShift   *WorkShiftModel `gorm:"foreignKey:WorkShiftId" json:"work_shift"`
}
