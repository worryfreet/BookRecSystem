package request

import (
	"BookRecSystem/model/system"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	ID          uint
	UUID        uuid.UUID
	UserID      uint
	BufferTime  int64
	AuthorityId uint
	jwt.StandardClaims
}

// User cache structure
type UserCache struct {
	UUID        string                `redis:"uuid"`
	ID          uint                  `redis:"id"`
	DeptId      uint                  `redis:"deptId"`
	AuthorityId []uint                `redis:"authorityId"`
	Authority   []system.SysAuthority `redis:"-"`
}

// User cache structure
type UserCacheRedis struct {
	ID          uint   `redis:"id"`
	DeptId      uint   `redis:"deptId"`
	AuthorityId []byte `redis:"authorityId"`
}
