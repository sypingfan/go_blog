package settings_api

import (
	"admin/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(res.SettingError, c)
}
