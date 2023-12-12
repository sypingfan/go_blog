/**
* Created by GoLand.
* User: 平凡
* Date: 2023/12/12
* Time: 22:21
* Remark: 登录表
 */
package models

// LoginDataModel 登录信息表 统计用户登录数据 id,用户id,用户昵称,用户token,登录设备,登录时间
type LoginDataModel struct {
	MODEL
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"` //登录的ip
	NickName  string    `gorm:"size:42" json:"nick_name"`
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"` // 登录设备
	Addr      string    `gorm:"size:64" json:"addr"`
}
