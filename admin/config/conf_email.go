/**
* Created by GoLand.
* User: 平凡
* Date: 2023/12/20
* Time: 22:45
* Remark: 邮箱配置
 */
package config

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"` // 发件人游戏
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"` // 默认的发件人名字
	UseSSL           bool   `json:"use_ssl" yaml:"use_ssl"`                       // 是否使用ssl
	UserTls          bool   `json:"user_tls" yaml:"user_tls"`
}
