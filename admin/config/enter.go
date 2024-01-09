package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Upload   Upload   `yaml:"upload"`
	QQ       QQ       `json:"qq"`
	QiNiu    QiNiu    `json:"qi_niu"`
	Email    Email    `json:"email"`
	Jwt      Jwt      `json:"jwt"`
}
