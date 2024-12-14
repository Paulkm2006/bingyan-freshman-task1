package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type IntArray []int

func (a *IntArray) Scan(value interface{}) error {
	if value == nil {
		*a = IntArray{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan IntArray")
	}

	return json.Unmarshal(bytes, a)
}

func (a IntArray) Value() (driver.Value, error) {
	if a == nil {
		return "[]", nil
	}
	return json.Marshal(a)
}

type Node struct {
	NID         int      `json:"nid" gorm:"primaryKey;autoIncrement;index" query:"nid"`
	Name        string   `json:"name" query:"name"`
	Description string   `json:"description" query:"description"`
	Article     int      `json:"article" query:"article"`
	Moderators  IntArray `json:"moderators" query:"moderators" gorm:"type:json"`
}
