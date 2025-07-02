package retailoutlet

type CreateRetailOutletDto struct {
	FullName   string `json:"fullName"  validate:"required,min=2"`
	Address    string `json:"address" validate:"required,min=5"`
	OpenedDate int64  `json:"openedDate" validate:"required"`
	ClosedDate int64  `json:"closedDate"`
	SeoId      int16  `json:"seoId" validate:"required"`
}
