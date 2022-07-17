package database

import (
	"time"
)

type User struct {
	CID               uint      `json:"cid" gorm:"primaryKey"`
	FirstName         string    `json:"firstname" gorm:"type:varchar(128)"`
	LastName          string    `json:"lastname" gorm:"type:varchar(128)"`
	Email             string    `json:"email" gorm:"type:varchar(128)"`
	OperatingInitials string    `json:"oi" gorm:"type:varchar(2)"`
	ControllerType    string    `json:"controllerType" gorm:"type:varchar(10)"`
	DelCertification  string    `json:"delCertification" gorm:"type:varchar(15)"`
	GndCertification  string    `json:"gndCertification" gorm:"type:varchar(15)"`
	LclCertification  string    `json:"lclCertification" gorm:"type:varchar(15)"`
	AppCertification  string    `json:"appCertification" gorm:"type:varchar(15)"`
	CtrCertification  string    `json:"ctrCertification" gorm:"type:varchar(15)"`
	RatingID          int       `json:"-"`
	Rating            Rating    `json:"rating"`
	Status            string    `json:"status" gorm:"type:varchar(10)"`
	Roles             []*Role   `json:"roles" gorm:"many2many:user_roles"`
	DiscordID         string    `json:"discord_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

var CertificationOptions = map[string]string{
	"none":      "none",
	"training":  "training",
	"solo":      "solo",
	"certified": "certified",
	"major":     "major",
	"cantrain":  "cantrain",
}

var ControllerTypeOptions = map[string]string{
	"none":    "none",
	"visitor": "visitor",
	"home":    "home",
}

var ControllerStatusOptions = map[string]string{
	"none":     "none",
	"active":   "active",
	"inactive": "inactive",
	"loa":      "loa",
}
