/**
* Created by GoLand.
* User: 平凡
* Date: 2024/1/12
* Time: 21:31
* Remark: 公共
 */
package common

import (
	"admin/global"
	"admin/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

// 分页
func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	count = DB.Select("id").Find(&list).RowsAffected

	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Limit(option.Limit).Offset(offset).Find(&list).Error

	return list, count, err

}
