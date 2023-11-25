package api

import "admin/api/settings_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}

var ApigroupApp = new(ApiGroup)
