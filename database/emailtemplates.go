package database

import "time"

type EmailTemplate struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(25)"`
	Subject   string    `json:"subject" gorm:"type:varchar(128)"`
	Body      string    `json:"body" gorm:"type:text"`
	EditGroup string    `json:"edit_group" gorm:"type:varchar(25)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
