/**
* Created by GoLand.
* User: 平凡
* Date: 2023/12/20
* Time: 22:45
* Remark: jwt配置 身份验证
 */
package config

type Jwt struct {
	Secret  string `json:"secret" yaml:"secret"`   // 秘钥
	Expires int    `json:"expires" yaml:"expires"` // 过期时间
	Issuer  string `json:"issuer" yaml:"issuer"`   // 颁发人
}
