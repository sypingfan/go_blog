/**
* Created by GoLand.
* User: 平凡
* Date: 2023/12/12
* Time: 22:00
* Remark: 执行db命令 迁移表结构
 */
package flag

import (
	"admin/global"
	"admin/models"
)

func Makemigrations() {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDb").
		AutoMigrate(
			&models.BannerModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.AdvertModel{},
			&models.UserModel{},
			&models.CommentModel{},
			&models.ArticleModel{},
			&models.MenuModel{},
			&models.MenuBannerModel{},
			&models.FadeBackModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		global.Log.Error("[ error ] 生成数据表结构失败")
	}
	global.Log.Info("[ success ] 生成数据表结构成功！")
}
