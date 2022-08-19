package request

type SysAuthorityCreateRequest struct {
	Level         uint   `json:"level"`
	AuthorityName string `json:"authorityName"` // 角色名
	DataScope     string `json:"dataScope"`
	DeptId        []uint `json:"deptId"`
}

type SysAuthorityUpdateRequest struct {
	AuthorityId   uint   `json:"authorityId"`
	Level         uint   `json:"level"`
	AuthorityName string `json:"authorityName"` // 角色名
	DataScope     string `json:"dataScope"`
	DeptId        []uint `json:"deptId"`
}

type SysAuthorityUserListRequest struct {
	AuthorityId uint   `json:"authorityId"`
	UserList    []uint `json:"userList"`
}
