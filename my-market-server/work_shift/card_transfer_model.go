package workshift

type CardTransferModel struct {
	ID          int `gorm:"primaryKey"`
	Sum         float32
	WorkShiftId int            `gorm:"column:work_shift_id"`
	WorkShift   WorkShiftModel `gorm:"foreignKey:WorkShiftId"`
}
