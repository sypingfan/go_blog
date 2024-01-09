package settings_api

import (
	"admin/global"
	"admin/models/res"
	"github.com/gin-gonic/gin"
)

type SettingUri struct {
	Name string `uri:"name"`
}

// SettingsInfoView 显示某一项的配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)

	}
}
