package database

import (
	"database/sql/driver"
	"time"
)

type User struct {
	CID               uint        `json:"cid" gorm:"primaryKey"`
	FirstName         string      `json:"firstname" gorm:"type:varchar(128)"`
	LastName          string      `json:"lastname" gorm:"type:varchar(128)"`
	Email             string      `json:"email" gorm:"type:varchar(128)"`
	OperatingInitials string      `json:"oi" gorm:"type:varchar(2)"`
	ControllerType    string      `json:"controllerType" gorm:"type:enum('none', 'visitor', 'home')"`
	DelCertification  CertOptions `json:"delCertification" gorm:"type:enum('none','training','solo','certified','major','cantrain')"`
	GndCertification  CertOptions `json:"gndCertification" gorm:"type:enum('none','training','solo','certified','major','cantrain')"`
	LclCertification  CertOptions `json:"lclCertification" gorm:"type:enum('none','training','solo','certified','major','cantrain')"`
	AppCertification  CertOptions `json:"appCertification" gorm:"type:enum('none','training','solo','certified','major','cantrain')"`
	CtrCertification  CertOptions `json:"ctrCertification" gorm:"type:enum('none','training','solo','certified','major','cantrain')"`
	RatingID          int         `json:"-"`
	Rating            Rating      `json:"rating"`
	Status            string      `json:"status" gorm:"type:enum('none', 'active', 'inactive', 'loa')"`
	Roles             []*Role     `json:"roles" gorm:"many2many:user_roles"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

type CertOptions string

const (
	NONE      CertOptions = "none"
	TRAINING  CertOptions = "training"
	SOLO      CertOptions = "solo"
	CERTIFIED CertOptions = "certified"
	MAJOR     CertOptions = "major"
	CAN_TRAIN CertOptions = "cantrain"
)

func (co *CertOptions) Scan(value interface{}) error {
	*co = CertOptions(value.(string))
	return nil
}

func (co CertOptions) Value() (driver.Value, error) {
	return string(co), nil
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
