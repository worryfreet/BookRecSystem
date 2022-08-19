package system

import (
	"time"

	"gorm.io/gorm"
)

type SysAuthority struct {
	AuthorityId   uint           `json:"authorityId" gorm:"not null;autoIncrement;primary_key;comment:角色ID;size:90"` // 角色ID
	Level         uint           `json:"level" gorm:"not null;comment:角色等级0最大"`
	AuthorityName string         `json:"authorityName" gorm:"not null;comment:角色名"` // 角色名
	DataScope     string         `json:"dataScope" gorm:"not null"`
	DefaultRouter string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
	CreateBy      uint           //创建人
	UpdateBy      uint           //更新人
	CreatedAt     time.Time      // 创建时间
	UpdatedAt     time.Time      // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"index;" json:"-"` // 删除时间
}
