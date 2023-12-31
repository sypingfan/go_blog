/**
* Created by GoLand.
* User: 平凡
* Date: 2023/12/20
* Time: 22:46
* Remark: qq的配置
 */
package config

import "fmt"

type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"` // 登录之后的回调地址
}

func (q QQ) GetPath() string {
	if q.Key == "" || q.AppID == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://qraph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&cilent_id=%s&redirect_url=%s", q.AppID, q.Redirect)
}
