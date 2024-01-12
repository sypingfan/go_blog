/**
* Created by GoLand.
* User: 平凡
* Date: 2024/1/12
* Time: 21:06
* Remark: 图片列表
 */
package images_api

import (
	"admin/models"
	"admin/models/res"
	"admin/service/common"
	"github.com/gin-gonic/gin"
)

// ImageListView 图片列表
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	res.OkWithList(list, count, c)
}
