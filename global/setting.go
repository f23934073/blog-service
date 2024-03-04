package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
