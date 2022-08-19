package system

type CasbinModel struct {
	PType       string `json:"p_type" gorm:"column:p_type"`
	AuthorityId string `json:"role_name" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}
