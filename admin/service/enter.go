/**
* Created by GoLand.
* User: 平凡
* Date: 2025/11/4
* Time: 20:34
* Remark: New挂载上传文件
 */
package service

import "admin/service/image_ser"

type ImagesGroup struct {
	ImageService image_ser.ImageService
}

var ServiceApp = new(ImagesGroup)
