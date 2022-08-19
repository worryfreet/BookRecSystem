package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

type PageInfoOrder struct {
	PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}

// GetById Find by id structure
type GetById struct {
	ID uint `json:"id" form:"id"` // 主键ID
}

type IdsReq struct {
	Ids []uint `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint // 角色ID
}

type Empty struct{}
