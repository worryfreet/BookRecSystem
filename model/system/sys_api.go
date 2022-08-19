package system

import "BookRecSystem/global"

type SysApi struct {
	global.GSD_MODEL
	Path        string `json:"path" gorm:"not null;comment:api路径"`             // api路径
	Description string `json:"description" gorm:"not null;comment:api中文描述"`    // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"not null;comment:api组"`          // api组
	Method      string `json:"method" gorm:"not null;default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}
