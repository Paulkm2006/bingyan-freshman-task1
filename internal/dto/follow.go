package dto

import "time"

type Follow struct {
	UID      int       `json:"uid" gorm:"primaryKey;index"`
	Followee int       `json:"followee" gorm:"primaryKey;index"`
	Created  time.Time `json:"created" gorm:"autoCreateTime"`
}
