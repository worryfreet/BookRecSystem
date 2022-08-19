package router

import (
	"BookRecSystem/router/book_rec"
	"BookRecSystem/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	BookRec book_rec.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
