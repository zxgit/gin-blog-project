package global

import (
	"github.com/zxgit/gin-blog-project/pkg/logger"
	"github.com/zxgit/gin-blog-project/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger *logger.Logger
)
