package dto

import "time"

type Comment struct {
	CID     int       `json:"cid" gorm:"primaryKey;autoIncrement;index" query:"cid"`
	UID     int       `json:"uid" gorm:"index" query:"uid"`
	Content string    `json:"content" query:"content"`
	PID     int       `json:"pid" gorm:"index" query:"pid"`
	Created time.Time `json:"created" gorm:"autoCreateTime" query:"created"`
}
