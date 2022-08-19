package system

import (
	"BookRecSystem/global"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.GSD_MODEL
	UserID      uint           `json:"user_id" gorm:"not null;comment:用户学工号"`
	Username    string         `json:"userName" gorm:"not null;comment:用户登录名"`
	Password    string         `json:"password"  gorm:"not null;comment:用户登录密码;type:longtext"`
	NickName    string         `json:"nickName" gorm:"not null;default:系统用户;comment:用户昵称"`
	HeaderImg   string         `json:"headerImg" gorm:"not null;default:uploads/file/head.jpg;comment:用户头像"`
	Email       string         `json:"email" gorm:"not null;comment:用户邮箱"`
	College     string         `json:"college" gorm:"not null; comment: 所属学院"`
	Grade       uint8          `json:"grade" gorm:"not null; comment: 所属年级"`
	Phone       string         `json:"phone" gorm:"comment:用户手机号"`
	Labels      string         `json:"labels" gorm:"comment: 用户读书兴趣"`
	AuthorityId uint           `json:"authorityId" gorm:"default:1;comment:用户角色ID"`
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	UUID        uuid.UUID      `json:"uuid" gorm:"not null;comment:用户UUID"`
}
