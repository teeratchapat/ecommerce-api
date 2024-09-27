package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name       string `json:"name"`
	Route      string `json:"route"`
	Icon       string `json:"icon"`
	IsChildren bool   `json:"is_children"`
	ParentID   *uint  `json:"parent_id"`
	Children   []Menu `json:"children" gorm:"foreignKey:ParentID"`
}
