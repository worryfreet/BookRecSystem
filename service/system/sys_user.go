package system

import (
	"BookRecSystem/global"
	commonReq "BookRecSystem/model/common/request"
	"BookRecSystem/model/system"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type UserService struct {
}

func (u UserService) UpdatePassword(user *system.SysUser, password string) (userReturn *system.SysUser, err error) {
	//password, err = utils.RsaPubEncode(password)
	if err != nil {
		return
	}
	err = global.GSD_DB.Where("user_id = ? AND password = ?", user.UserID, user.Password).First(&user).Update("password", password).Error
	return
}

func (u UserService) GetUserInfo(id uint) (user system.SysUser, err error) {
	err = global.GSD_DB.Where("id = ?", id).Find(&user).Error
	return
}

func (u UserService) Login(user *system.SysUser) (*system.SysUser, error) {
	var userReturn system.SysUser
	err := global.GSD_DB.Where("user_id = ? AND password = ?", user.UserID, user.Password).First(&userReturn).Error
	return &userReturn, err
}

func (u UserService) Register(user system.SysUser) error {
	var uTmp system.SysUser
	if !errors.Is(global.GSD_DB.Where("user_id = ?", user.UserID).First(&uTmp).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户已经被注册")
	}
	global.GSD_DB.Create(&user)
	return nil
}

func (u UserService) FindUserById(userId uint) (user system.SysUser, err error) {
	err = global.GSD_DB.Where("user_id = ?", userId).Find(&user).Error
	return
}

func (u UserService) UpdateUserInterest(id uint, labels []string) (err error) {
	labelsBytes, err := json.Marshal(labels)
	if err != nil {
		return
	}
	err = global.GSD_DB.Where("id = ?", id).Update("label", string(labelsBytes)).Error
	return
}

func (u UserService) GetUserList(pageInfo commonReq.PageInfo) (userList []system.SysUser, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	err = global.GSD_DB.Limit(limit).Offset(offset).Find(&userList).Error
	return
}
