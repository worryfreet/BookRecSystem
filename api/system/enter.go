package system

import "BookRecSystem/service"

type ApiGroup struct {
	UserApi
}

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
