package service

import (
	"BookRecSystem/service/book_rec"
	"BookRecSystem/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.SysGroup
	BookRecGroup       book_rec.BookRecGroup
}

var ServiceGroupApp = new(ServiceGroup)
