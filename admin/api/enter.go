package api

import (
	"admin/api/images_api"
	"admin/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
}

var ApigroupApp = new(ApiGroup)
