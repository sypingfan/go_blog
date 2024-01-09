package routers

import (
	"admin/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApigroupApp.SettingsApi
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)

}
