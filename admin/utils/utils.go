/**
* Created by GoLand.
* User: 平凡
* Date: 2024/1/10
* Time: 21:39
* Remark: 公共方法
 */
package utils

// InList 判断key是否存在与列表中
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}
