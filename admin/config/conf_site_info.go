/**
* Created by GoLand.
* User: 平凡
* Date: 2023/12/13
* Time: 22:05
* Remark: 站点配置-网站信息
 */
package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at"`
	BeiAn       string `yaml:"bei_an"`
	Title       string `yaml:"title"`
	QQImage     string `yaml:"qq_image"`
	Version     string `yaml:"version"`
	Email       string `yaml:"email"`
	WechatImage string `yaml:"wechat_image"`
	Name        string `yaml:"name"`
	Job         string `yaml:"job"`
	Addr        string `yaml:"addr"`
	Slogan      string `yaml:"slogan"`
	SloganEN    string `yaml:"slogan_en"`
	Web         string `yaml:"web"`
	BiliBiliURL string `yaml:"bilibili_url"`
	GiteeUrl    string `yaml:"gitee_url"`
	GithubUrl   string `yaml:"github_url"`
}
