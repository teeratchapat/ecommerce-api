package seed

import (
	"menu-api/models"

	"gorm.io/gorm"
)

func SeedMenus(DB *gorm.DB) {
	var count int64
	DB.Model(&models.Menu{}).Count(&count)

	if count == 0 {
		menu1 := models.Menu{
			Name:       "ข้อมูลสรุป",
			Route:      "/dashboards",
			Icon:       "fi-rr-chart-pie-alt",
			IsChildren: false,
		}
		DB.Create(&menu1)

		child1 := models.Menu{
			Name:       "สินทรัพย์",
			Route:      "/dashboards/asset",
			ParentID:   &menu1.ID,
			IsChildren: true,
		}
		child2 := models.Menu{
			Name:       "สภาพทาง",
			Route:      "/dashboards/condition",
			ParentID:   &menu1.ID,
			IsChildren: true,
		}

		DB.Create(&child1)
		DB.Create(&child2)
	}
}
