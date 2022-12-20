package global

import (
	"go_im/common/config"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Db     *gorm.DB
)
