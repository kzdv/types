package database

import "time"

type EmailJob struct {
	ID            int       `json:"id"`
	SendTo        string    `json:"send_to" gorm:"type:varchar(128)"`
	EmailTemplate string    `json:"email_template" gorm:"type:varchar(25)"`
	Variables     string    `json:"variables" gorm:"type:text"`
	Completed     bool      `json:"completed"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
