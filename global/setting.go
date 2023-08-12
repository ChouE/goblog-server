package global

import (
	"github.com/ChouE/goblog-server/pkg/logger"
	"github.com/ChouE/goblog-server/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
