package retailoutlet

type CreateRetailOutletDto struct {
	FullName   string `json:"full_name"  validate:"required,min=2"`
	Address    string `json:"address" validate:"required,min=5"`
	OpenedDate int64  `json:"opened_date" validate:"required"`
	ClosedDate int64  `json:"closed_date" validate:"required"`
	SeoId      int16  `json:"seo_id" validate:"required"`
}
