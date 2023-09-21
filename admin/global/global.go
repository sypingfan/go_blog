package global

import (
	"admin/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
