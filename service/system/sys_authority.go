package system

import (
	"BookRecSystem/global"
	"BookRecSystem/model/common/request"
	"BookRecSystem/model/system"
	"errors"
	"gorm.io/gorm/clause"
	"strconv"

	"gorm.io/gorm"
)

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

// CreateAuthority
// @description: 创建一个角色
// @param: auth model.SysAuthority
// @return: err error, authority model.SysAuthority
func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (err error, authority system.SysAuthority) {
	err = global.GSD_DB.Create(&auth).Error
	return err, auth
}

// DeleteAuthority
// @description: 删除角色
// @param: auth *model.SysAuthority
// @return: err error
func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) (err error) {
	if !errors.Is(global.GSD_DB.Where("sys_authority_authority_id = ?", auth.AuthorityId).First(&system.SysUseAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	return global.GSD_DB.Transaction(func(tx *gorm.DB) error {
		db := global.GSD_DB.Where("authority_id = ?", auth.AuthorityId).First(auth)
		if err := db.Delete(auth).Error; err != nil {
			return err
		}
		if success := CasbinServiceApp.ClearCasbin(0, strconv.Itoa(int(auth.AuthorityId))); !success {
			return nil
		}
		return nil
	})
}

// GetAuthorityInfoList
// @description: 分页获取数据
// @param: info request.PageInfo
// @return: err error, list interface{}, total int64
func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var authority []system.SysAuthority
	err = global.GSD_DB.Limit(limit).Offset(offset).Find(&authority).Error
	return err, authority, total
}

// GetAuthorityInfo
// @description: 获取所有角色信息
// @param: auth model.SysAuthority
// @return: err error, sa model.SysAuthority
func (authorityService *AuthorityService) GetAuthorityInfo(auth system.SysAuthority) (err error, sa system.SysAuthority) {
	err = global.GSD_DB.Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

// GetAuthorityInfoByIDs
// @description: 根据角色id切片获取所有角色信息
// @param: authIds []uint
// @return: err error, sa []model.SysAuthority
func (authorityService *AuthorityService) GetAuthorityInfoByIDs(authIds []uint) (err error, sa []system.SysAuthority) {
	err = global.GSD_DB.Where("authority_id in (?)", authIds).Find(&sa).Error
	return err, sa
}

// GetAuthorityBasicInfo
// @description: 获取基本角色信息
// @param: auth model.SysAuthority
// @return: err error, sa model.SysAuthority
func (authorityService *AuthorityService) GetAuthorityBasicInfo(auth system.SysAuthority) (err error, sa system.SysAuthority) {
	err = global.GSD_DB.Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

// AddUserList
// @description: 角色下新增用户
// @param: authUserList systemReq.SysAuthorityUserListRequest
// @return: error
func (authorityService *AuthorityService) AddUserList(authUserList []system.SysUseAuthority) error {
	return global.GSD_DB.Model(&system.SysUseAuthority{}).Clauses(clause.OnConflict{DoNothing: true}).Create(&authUserList).Error
}

// DeleteUserList
// @description: 角色下新增用户
// @param: authUserList systemReq.SysAuthorityUserListRequest
// @return: error
func (authorityService *AuthorityService) DeleteUserList(authId uint, userIdList []uint) error {
	return global.GSD_DB.Where("`sys_user_id`in (?) AND `sys_authority_authority_id` = ?", userIdList, authId).Delete(&system.SysUseAuthority{}).Error
}
