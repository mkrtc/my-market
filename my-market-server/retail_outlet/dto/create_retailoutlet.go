package retailoutlet_dto

type CreateRetailOutletDto struct {
	FullName   string `json:"full_name"  validate:"required"`
	Address    string `json:"address"`
	OpenedDate int64  `json:"opened_date"`
	ClosedDate int64  `json:"closed_date"`
}
