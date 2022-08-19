package system

import (
	"BookRecSystem/global"
	"BookRecSystem/model/system/request"
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type JwtService struct {
}

var authorityService = AuthorityService{}

// GetRedisJWT
// @description: 从redis取jwt
// @param: userName string
// @return: err error, redisJWT string
func (jwtService *JwtService) GetRedisJWT(userId uint) (redisJWT string, err error) {
	redisJWT, err = global.GSD_REDIS.Get(context.Background(), strconv.Itoa(int(userId))).Result()
	return redisJWT, err
}

// SetRedisJWT
// @description: jwt存入redis并设置过期时间
// @param: jwt string, userName string
// @return: err error
func (jwtService *JwtService) SetRedisJWT(jwt string, userId uint) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GSD_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GSD_REDIS.Set(context.Background(), strconv.Itoa(int(userId)), jwt, timer).Err()
	return err
}

// SetRedisUserInfo
// @description: 从redis获取用户信息
// @param: uuid string
// @return: err error
func (jwtService *JwtService) SetRedisUserInfo(uuid string, userInfo request.UserCache) (err error) {
	_, err = global.GSD_REDIS.Pipelined(context.Background(), func(rdb redis.Pipeliner) error {
		rdb.HSet(context.Background(), uuid, "id", userInfo.ID)
		rdb.HSet(context.Background(), uuid, "uuid", userInfo.UUID)
		rdb.HSet(context.Background(), uuid, "deptId", userInfo.DeptId)
		authorityJson, _ := json.Marshal(userInfo.AuthorityId)
		rdb.HSet(context.Background(), uuid, "authorityId", authorityJson)
		return nil
	})
	return err
}

// GetRedisUserInfo
// @description: 从redis获取用户信息
// @param: jwt string, userName string
// @return: err error
func (jwtService *JwtService) GetRedisUserInfo(uuid string) (redisUserInfo request.UserCache, err error) {
	var userInfoRedis request.UserCacheRedis
	err = global.GSD_REDIS.HGetAll(context.Background(), uuid).Scan(&userInfoRedis)
	if err != nil || userInfoRedis.ID == 0 {
		return redisUserInfo, errors.New("查询用户缓存失败")
	}
	redisUserInfo.UUID = uuid
	redisUserInfo.ID = userInfoRedis.ID
	if err != nil || userInfoRedis.ID == 0 {
		return redisUserInfo, errors.New("查询用户缓存失败")
	}
	return
}

// DelRedisUserInfo
// @description: 从redis删除用户信息
// @param: uuid string
// @return: err error
func (jwtService *JwtService) DelRedisUserInfo(uuid string) (err error) {
	return global.GSD_REDIS.HDel(context.Background(), uuid).Err()
}

// DelRedisUserInfoList
// @description: 从redis删除用户信息
// @param: uuid string
// @return: err error
func (jwtService *JwtService) DelRedisUserInfoList(uuid []string) (err error) {
	_, err = global.GSD_REDIS.Pipelined(context.Background(), func(pipeliner redis.Pipeliner) error {
		for _, uid := range uuid {
			err = pipeliner.HDel(context.Background(), uid).Err()
		}
		return err
	})
	return
}

// UpdateUserListAuthorities
// @description: 从redis更新用户角色
// @param: userAuth map[string][]uint
// @return: err error
func (jwtService *JwtService) UpdateUserListAuthorities(userAuth map[string][]uint) (err error) {
	_, err = global.GSD_REDIS.Pipelined(context.Background(), func(pipeliner redis.Pipeliner) error {
		for uuid, authorityIds := range userAuth {
			authorityIdJson, _ := json.Marshal(authorityIds)
			err = pipeliner.HSet(context.Background(), uuid, "authorityId", authorityIdJson).Err()
		}
		return err
	})
	return
}

// UpdateUserAuthorities
// @description: 从redis更新用户角色
// @param: uuid string, authorityIds []uint
// @return: err error
func (jwtService *JwtService) UpdateUserAuthorities(uuid string, authorityIds []uint) (err error) {
	authorityIdJson, _ := json.Marshal(authorityIds)
	err = global.GSD_REDIS.HSet(context.Background(), uuid, "authorityId", authorityIdJson).Err()
	if err != nil {
		return err
	}
	return
}
