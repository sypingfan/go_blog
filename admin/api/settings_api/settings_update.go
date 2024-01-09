/**
* Created by GoLand.
* User: 平凡
* Date: 2023/12/13
* Time: 22:21
* Remark: 修改站点信息
 */
package settings_api

import (
	"admin/config"
	"admin/core"
	"admin/global"
	"admin/models/res"
	"github.com/gin-gonic/gin"
)

// SettingsInfoUpdateView 修改某一项的配置信息
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwt
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwt = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}
	core.SetYaml()
	res.OkWith(c)
}
