/**
* Created by GoLand.
* User: 平凡
* Date: 2024/1/10
* Time: 21:49
* Remark: md5加密
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 md5加密
func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
