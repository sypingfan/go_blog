package api

import (
	"admin/api/advert_api"
	"admin/api/images_api"
	"admin/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
}

var ApigroupApp = new(ApiGroup)
