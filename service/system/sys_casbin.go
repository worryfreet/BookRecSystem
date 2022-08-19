package system

import (
	"BookRecSystem/global"
	"BookRecSystem/model/system"
	"BookRecSystem/model/system/request"
	"errors"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type CasbinService struct {
}

var CasbinServiceApp = new(CasbinService)

// UpdateCasbin
// @description: 更新casbin权限
// @param: authorityId string, casbinInfos []request.CasbinInfo
// @return: error
func (casbinService *CasbinService) UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	casbinService.ClearCasbin(0, authorityId)
	var rules [][]string
	for _, v := range casbinInfos {
		cm := system.CasbinModel{
			PType:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	success, _ := global.GSD_Casbin.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

//@function: UpdateCasbinApi
//@description: API更新随动
//@param: oldPath string, newPath string, oldMethod string, newMethod string
//@return: error

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.GSD_DB.Table("casbin_rule").Model(&system.CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

//@function: UpdateUserAuthority
//@description: 用户角色更新
//@param: userId uint 用户id, authorityIds []uint 角色id
//@return: error

func (casbinService *CasbinService) UpdateUserAuthority(userId uint, authorityIds []uint) error {
	success, err := global.GSD_Casbin.DeleteRolesForUser(strconv.Itoa(int(userId)))
	var roles []string
	for _, authorityId := range authorityIds {
		roles = append(roles, strconv.Itoa(int(authorityId)))
	}
	success, _ = global.GSD_Casbin.AddRolesForUser(strconv.Itoa(int(userId)), roles)
	if !success {
		return err
	}
	return nil
}

//@function: DeleteUserAuthority
//@description: 用户角色更新
//@param: userId uint 用户id
//@return: error

func (casbinService *CasbinService) DeleteUserAuthority(userId uint) error {
	success, err := global.GSD_Casbin.DeleteRolesForUser(strconv.Itoa(int(userId)))
	if !success {
		return err
	}
	return nil
}

//@function: GetPolicyPathByAuthorityId
//@description: 获取权限列表
//@param: authorityId string
//@return: pathMaps []request.CasbinInfo

func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	list := global.GSD_Casbin.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

//@function: ClearCasbin
//@description: 清除匹配的权限
//@param: v int, p ...string
//@return: bool

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	success, _ := global.GSD_Casbin.RemoveFilteredPolicy(v, p...)
	return success
}
