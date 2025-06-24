package workshift

type CardTransferModel struct {
	ID          int `gorm:"primaryKey; autoIncrement"`
	Sum         float32
	WorkShiftId int            `gorm:"column:work_shift_id"`
	WorkShift   WorkShiftModel `gorm:"foreignKey:WorkShiftId"`
}
