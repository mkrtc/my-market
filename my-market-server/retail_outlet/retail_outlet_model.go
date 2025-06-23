package retailoutlet

import (
	"time"
)

type RetailOutletModel struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	FullName   string    `gorm:"column:full_name" json:"full_name"`
	Address    string    `json:"address"`
	OpenedDate time.Time `gorm:"column:opened_date" json:"opened_date"`
	ClosedDate time.Time `gorm:"column:closed_date" json:"closed_date"`
	SeoId      int       `form:"column:seo_id" json:"seo_id"`
	Seo        SeoModel  `gorm:"foreignKey:SeoId" json:"seo"`
}
