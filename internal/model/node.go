package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm/clause"
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

func CreateNode(node *Node) (*Node, error) {
	result := db.Model(&Node{}).Clauses(clause.Returning{}).Create(node)
	if result.Error != nil {
		return nil, result.Error
	}
	return node, nil
}

func GetNodes() ([]Node, error) {
	var nodes []Node
	result := db.Find(&nodes)
	if result.Error != nil {
		return nil, result.Error
	}
	return nodes, nil
}

func GetNodeByNID(nid int) (*Node, error) {
	var node Node
	result := db.Where("n_id = ?", nid).First(&node)
	if result.Error != nil {
		return nil, result.Error
	}
	return &node, nil
}

func AddModerator(uid int, nid int) error {
	node, err := GetNodeByNID(nid)
	if err != nil {
		return err
	}
	node.Moderators = append(node.Moderators, uid)
	result := db.Save(node)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteModerator(uid int, nid int) error {
	node, err := GetNodeByNID(nid)
	if err != nil {
		return err
	}
	for i, v := range node.Moderators {
		if v == uid {
			node.Moderators = append(node.Moderators[:i], node.Moderators[i+1:]...)
			break
		}
	}
	result := db.Save(node)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteNode(nid int) error {
	result := db.Where("nid = ?", nid).Delete(&Node{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func IncrArticle(nid int) error {
	node, err := GetNodeByNID(nid)
	if err != nil {
		return err
	}
	node.Article++
	result := db.Save(node)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DecrArticle(nid int) error {
	node, err := GetNodeByNID(nid)
	if err != nil {
		return err
	}
	node.Article--
	result := db.Save(node)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
