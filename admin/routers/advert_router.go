/**
* Created by GoLand.
* User: 平凡
* Date: 2025/11/11
* Time: 20:37
* Remark: 广告路由
 */
package routers

import "admin/api"

func (router RouterGroup) AdvertRouter() {
	app := api.ApigroupApp.AdvertApi
	router.POST("adverts", app.AdvertCreateView)
}
