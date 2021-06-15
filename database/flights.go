package database

type Flights struct {
	Callsign    string  `json:"callsign" gorm:"primaryKey;type:varchar(10)"`
	CID         int     `json:"cid" gorm:"index"`
	Latitude    float32 `json:"latitude" gorm:"type:float(10,8)"`
	Longitude   float32 `json:"longitude" gorm:"type:float(11,8)"`
	Groundspeed int     `json:"groundspeed"`
	Heading     int     `json:"heading"`
	Altitude    int     `json:"altitude"`
	Aircraft    string  `json:"aircraft" gorm:"type:varchar(6)"`
	Route       string  `json:"route" gorm:"type:text"`
}
