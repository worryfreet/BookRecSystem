package system

import (
	"BookRecSystem/global"
	"BookRecSystem/middleware"
	"BookRecSystem/model/common/request"
	"BookRecSystem/model/common/response"
	"BookRecSystem/model/system"
	systemReq "BookRecSystem/model/system/request"
	systemRes "BookRecSystem/model/system/response"
	"BookRecSystem/utils"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

type UserApi struct {
}

// @Tags User
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户学工号, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /api/user/login [post]
func (b *UserApi) Login(c *gin.Context) {
	var l systemReq.Login
	_ = c.ShouldBindJSON(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//l.Password = utils.RsaPriDecode(l.Password)
	//私钥解密
	u := &system.SysUser{UserID: l.UserID, Password: l.Password}
	if user, err := userService.Login(u); err != nil {
		global.GSD_LOG.Error("登录失败! 用户名不存在或密码错误!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("用户名不存在或密码错误", c)
	} else {
		b.tokenNext(c, *user)
	}
}

// 登录以后签发jwt
func (b *UserApi) tokenNext(c *gin.Context, user system.SysUser) {
	j := &middleware.JWT{SigningKey: []byte(global.GSD_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := systemReq.CustomClaims{
		ID:         user.ID,
		UUID:       user.UUID,
		UserID:     user.UserID,
		BufferTime: global.GSD_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GSD_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "gsdPlus",                                             // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GSD_LOG.Error("获取token失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("获取token失败", c)
		return
	}
	userCache := systemReq.UserCache{
		ID:        user.ID,
		UUID:      user.UUID.String(),
		Authority: user.Authorities,
	}
	_ = jwtService.SetRedisUserInfo(user.UUID.String(), userCache)
	if !global.GSD_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	// 设置redisKey
	fmt.Println("准备redisKey")
	if _, err = jwtService.GetRedisJWT(user.UserID); err != nil && err != redis.Nil {
		global.GSD_LOG.Error("设置登录状态失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		fmt.Println("设置redisKey")
		if err := jwtService.SetRedisJWT(token, user.UserID); err != nil {
			global.GSD_LOG.Error("设置登录状态失败!", zap.Any("err", err), utils.GetRequestID(c))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		//设置用户缓存
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// @Tags User
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body systemReq.Register true "用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /api/user/register [post]
func (b *UserApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//私钥解密
	//r.Password = utils.RsaPriDecode(r.Password)
	user := &system.SysUser{
		UserID:    r.UserID,
		Username:  r.Username,
		Password:  r.Password,
		NickName:  r.NickName,
		HeaderImg: r.HeaderImg,
		Email:     r.Email,
		College:   r.College,
		Grade:     r.Grade,
		Phone:     r.Phone,
		UUID:      uuid.NewV4(),
	}
	labelBytes, _ := json.Marshal(r.Labels)
	user.Labels = string(labelBytes)
	err := userService.Register(*user)
	if err != nil {
		global.GSD_LOG.Error("注册失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("注册失败", c)
	} else {
		response.Ok(c)
	}
}

// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/user/userList [get]
func (b *UserApi) UserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userList, err := userService.GetUserList(pageInfo)
	if err != nil {
		global.GSD_LOG.Error("获取失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(userList, c)
}

// @Tags User
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body systemReq.UpdatePwd true "用户学工号, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /api/user/updatePassword [put]
func (b *UserApi) UpdatePassword(c *gin.Context) {
	var user systemReq.UpdatePwd
	_ = c.ShouldBindJSON(&user)
	if err := utils.Verify(user, utils.UpdatePwdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//私钥解密
	//user.Password = utils.RsaPriDecode(user.Password)
	//user.NewPassword = utils.RsaPriDecode(user.NewPassword)
	u := &system.SysUser{
		UserID:   user.UserID,
		Password: user.Password,
	}
	if _, err := userService.UpdatePassword(u, user.NewPassword); err != nil {
		global.GSD_LOG.Error("修改失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("修改失败， 原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/user/userInfo [get]
func (b *UserApi) UserInfo(c *gin.Context) {
	id := utils.GetUserID(c)
	//fmt.Println("id: ---------------------------------", id)
	//uuid := utils.GetUserUuid(c)
	if userInfo, err := userService.GetUserInfo(id); err != nil {
		global.GSD_LOG.Error("获取用户信息失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取用户信息失败", c)
		return
	} else {
		response.OkWithDetailed(gin.H{"userInfo": userInfo}, "获取用户信息成功", c)
	}
}

// @Tags User
// @Summary 设置当前用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body []string true "读书兴趣"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /api/user/updateUserInterest [put]
func (b *UserApi) UpdateUserInterest(c *gin.Context) {
	var labels []string
	_ = c.ShouldBindJSON(&labels)
	if err := userService.UpdateUserInterest(utils.GetUserID(c), labels); err != nil {
		global.GSD_LOG.Error("设置失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("设置失败", c)
	} else {
		response.Ok(c)
	}
}
