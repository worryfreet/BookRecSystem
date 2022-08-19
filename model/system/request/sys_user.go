package request

import (
	commonReq "BookRecSystem/model/common/request"

	uuid "github.com/satori/go.uuid"
)

type Register struct {
	UserID    uint     `json:"user_id"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	NickName  string   `json:"nickName" gorm:"default:'GSDUser'"`
	HeaderImg string   `json:"header_img"`
	Phone     string   `json:"phone"`
	College   string   `json:"college"`
	Grade     uint8    `json:"grade"`
	Email     string   `json:"email"`
	Labels    []string `json:"labels"`
}

type Login struct {
	UserID   uint   `json:"user_id"`  // 用户学工号
	Password string `json:"password"` // 密码
}

type UpdatePwd struct {
	UserID      uint   `json:"user_id"`     // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint   `json:"id"`
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

// Modify  user's auth structure
type GetUserListOrder struct {
	SearchQuery string `json:"searchQuery"`
	commonReq.PageInfoOrder
}

// Modify  user's Info structure
type SetUserInfo struct {
	ID           uint      `json:"ID"`           //用户id
	DeptId       uint      `json:"deptId"`       //部门id
	HeadImg      string    `json:"headerImg"`    //头像
	NickName     string    `json:"nickName"`     //昵称
	Phone        string    `json:"phone"`        //手机号
	Email        string    `json:"email"`        //邮箱
	UUID         uuid.UUID `json:"uuid"`         //uuid
	AuthorityIds []uint    `json:"authorityIds"` //角色ID
}

type AuthorityUserName struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
}
