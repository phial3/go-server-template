package global

import (
	"github.com/jinzhu/gorm"
	"github.com/phial3/go-server-template/pkg/logger"
	"github.com/phial3/go-server-template/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
