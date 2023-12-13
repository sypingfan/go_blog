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

func (SettingsApi) SettingInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
