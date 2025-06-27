package retailoutlet

type CreateSeoDto struct {
	FullName  string `json:"fullName" validate:"required,min=5"`
	ShortName string `json:"shortName" validate:"required,min=3"`
	OrgName   string `json:"orgName" validate:"required,min=5"`
}
