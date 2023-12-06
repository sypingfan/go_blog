package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1 //管理员
	PermissionUser        Role = 2 //普通管理员
	PermissionVisitor     Role = 3 //游客
	PermissionDisableUser Role = 4 //被禁用的用户
)

func (s Role) MarshalJson() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "被禁言"
	default:
		str = "其他"
	}
	return str
}
