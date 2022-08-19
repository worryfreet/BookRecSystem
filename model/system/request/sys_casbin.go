package request

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId uint         `json:"authorityId"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/api/menu/getMenu", Method: "POST"},
		{Path: "/api/jwt/jsonInBlacklist", Method: "POST"},
		{Path: "/api/base/login", Method: "POST"},
		{Path: "/api/user/register", Method: "POST"},
		{Path: "/api/user/changePassword", Method: "POST"},
		{Path: "/api/user/setUserAuthority", Method: "POST"},
		{Path: "/api/user/setUserInfo", Method: "PUT"},
		{Path: "/api/user/getUserInfo", Method: "GET"},
		{Path: "/api/department/lists", Method: "POST"},
	}
}
