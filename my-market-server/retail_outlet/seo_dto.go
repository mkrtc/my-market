package retailoutlet

type CreateSeoDto struct {
	FullName  string `json:"full_name" validate:"required,min=5"`
	ShortName string `json:"short_name" validate:"required,min=3"`
	OrgName   string `json:"org_name" validate:"required,min=5"`
}
