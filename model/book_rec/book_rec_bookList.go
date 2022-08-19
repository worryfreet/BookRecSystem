package book_rec

import "BookRecSystem/model/system"

type BookList struct {
	UserID uint           `json:"user_id" gorm:"comment:用户ID;not null;"`
	User   system.SysUser `json:"user" gorm:"foreignKey:UserID;references:UserID;comment:用户"`
	BookID uint           `json:"book_id" gorm:"comment:图书ID;not null;"`
	Book   Book           `json:"book" gorm:"foreignKey:BookID;references:ID;comment:图书"`
}
