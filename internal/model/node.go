package model

import (
	"bingyan-freshman-task0/internal/dto"

	"gorm.io/gorm/clause"
)

func CreateNode(node *dto.Node) (*dto.Node, error) {
	result := db.Model(&dto.Node{}).Clauses(clause.Returning{}).Create(node)
	if result.Error != nil {
		return nil, result.Error
	}
	return node, nil
}

func GetNodes() ([]dto.Node, error) {
	var nodes []dto.Node
	result := db.Find(&nodes)
	if result.Error != nil {
		return nil, result.Error
	}
	return nodes, nil
}

func GetNodeByNID(nid int) (*dto.Node, error) {
	var node dto.Node
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
	result := db.Where("nid = ?", nid).Delete(&dto.Node{})
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
