package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name       string `json:"name"`                                // ชื่อของเมนู
	Route      string `json:"route"`                               // เส้นทางของเมนู
	Icon       string `json:"icon"`                                // ไอคอนของเมนู
	IsChildren bool   `json:"is_children"`                         // บ่งบอกว่าเป็นเมนูย่อยหรือไม่
	ParentID   *uint  `json:"parent_id"`                           // ถ้าเป็นเมนูย่อยจะอ้างถึงเมนูหลัก
	Children   []Menu `json:"children" gorm:"foreignKey:ParentID"` // เมนูย่อย (ถ้ามี)
}
