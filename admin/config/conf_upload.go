/**
* Created by GoLand.
* User: 平凡
* Date: 2024/1/7
* Time: 16:22
* Remark: 图片上传配置
 */
package config

type Upload struct {
	Size int    `yaml:"size" json:"size"` // 图片上传的大小
	Path string `yaml:"path" yaml:"path"` // 图片上传的目录
}
