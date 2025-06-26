package workshift

type CreateWorkShiftDto struct {
	Cash           float32                          `json:"cash" validate:"gte=0"`
	Cashless       float32                          `json:"cash_less" validate:"gte=0"`
	CashRegister   float32                          `json:"cash_register" validate:"gte=0"`
	RetailOutletId int                              `json:"retail_outlet_id" validate:"gte=0"`
	CardTransfers  *[]float32                       `json:"card_transfers" validate:"omitempty,dive,gte=0"`
	Expenses       *[]CreateExpenseWithWorkShiftDto `json:"expenses" validate:"omitempty,dive"`
}

type CreateExpenseWithWorkShiftDto struct {
	Article string  `json:"article" validate:"required"`
	Debit   float32 `json:"debit" validate:"gte=0"`
	Credit  float32 `json:"credit" validate:"gte=0"`
	Payed   bool    `json:"payed" validate:"required"`
}

type CreateCardTransferDto struct {
	Sum         float32 `json:"sum"`
	WorkShiftId int     `json:"work_shift_id"`
}

type CreateExpenseDto struct {
	Article     string  `json:"article" validate:"required"`
	Debit       float32 `json:"debit" validate:"required"`
	Credit      float32 `json:"credit" validate:"required"`
	Payed       bool    `json:"payed" validate:"required"`
	WorkShiftId int     `json:"work_shift_id" validate:"required"`
}
