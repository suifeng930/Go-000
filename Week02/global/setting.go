package global

import (
	"github.com/Go-000/Week02/pkg/setting"
	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	DataBaseSetting *setting.DatabaseSettingS
	DBEngine        *gorm.DB
)
