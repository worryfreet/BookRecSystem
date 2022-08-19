package book_rec

import "BookRecSystem/service"

type ApiGroup struct {
	BookRecApi
}

var bookRecService = service.ServiceGroupApp.BookRecGroup.BookRecService
