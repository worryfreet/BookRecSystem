package global

import (
	"time"

	"gorm.io/gorm"
)

type GSD_MODEL struct {
	ID        uint           `gorm:"primarykey;autoIncrement"` // 主键ID
	CreateBy  uint           //创建人
	UpdateBy  uint           //更新人
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"-"` // 删除时间
}
