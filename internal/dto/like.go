package dto

import "time"

type Like struct {
	LID     int       `json:"lid" gorm:"primaryKey;autoIncrement;index" query:"lid"`
	UID     int       `json:"uid" gorm:"index" query:"uid"`
	PID     int       `json:"pid" gorm:"index" query:"pid"`
	Created time.Time `json:"created" gorm:"autoCreateTime" query:"created"`
}
