package entities

type Role struct {
	ID          string `json:"id" gorm:"type:varchar(255);primaryKey"`
	Name        string `json:"name" gorm:"type:varchar(255);unique;not null"`
	Description string `json:"description" gorm:"type:text;not null"`
}
