package routers

import (
	"admin/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApigroupApp.SettingsApi
	router.GET("settings/:name", settingsApi.SettingInfoView)
	router.PUT("settings/:name", settingsApi.SettingInfoUpdateView)

}
