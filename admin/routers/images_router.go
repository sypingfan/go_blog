/**
* Created by GoLand.
* User: 平凡
* Date: 2024/1/7
* Time: 16:06
* Remark: 上传图片路由
 */
package routers

import "admin/api"

func (router RouterGroup) ImagesRouter() {
	app := api.ApigroupApp.ImagesApi
	router.POST("image", app.ImageUploadView)

}
