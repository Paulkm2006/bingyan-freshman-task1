package dto

import "time"

type Post struct {
	PID      int       `json:"pid" gorm:"primaryKey;autoIncrement;index" query:"pid"`
	Created  time.Time `json:"created" gorm:"autoCreateTime" query:"created"`
	UID      int       `json:"uid" gorm:"index" query:"uid"`
	Title    string    `json:"title" query:"title"`
	NID      int       `json:"nid" gorm:"index" query:"nid"`
	Likes    int       `json:"likes" gorm:"default:0"`
	Comments int       `json:"comments" gorm:"default:0"`
	Content  *string   `json:"content,omitempty" gorm:"-"` // Exclude from database
}

type Body struct {
	PID     int    `json:"pid" gorm:"primaryKey;index" query:"pid"`
	Content string `json:"content" query:"content"`
}
