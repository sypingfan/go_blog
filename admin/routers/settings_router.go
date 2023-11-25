package routers

import (
	"admin/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApigroupApp.SettingsApi
	router.GET("settings", settingsApi.SettingsInfoView)
}
