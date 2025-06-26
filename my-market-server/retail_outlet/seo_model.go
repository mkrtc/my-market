package retailoutlet

type SeoModel struct {
	ID            int                  `gorm:"primaryKey; autoIncrement" json:"id"`
	FullName      string               `gorm:"column:full_name" json:"full_name"`
	ShortName     string               `gorm:"column:short_name" json:"short_name"`
	OrgName       string               `gorm:"column:org_name" json:"org_name"`
	RetailOutlets *[]RetailOutletModel `gorm:"foreignKey:SeoId" json:"retail_otlets"`
}
