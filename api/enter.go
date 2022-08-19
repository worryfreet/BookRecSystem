package v1

import (
	"BookRecSystem/api/book_rec"
	"BookRecSystem/api/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	BookRecApiGroup book_rec.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
