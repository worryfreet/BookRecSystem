package book_rec

import (
	"BookRecSystem/global"
)

type Book struct {
	global.GSD_MODEL
	BookName    string `json:"bookName" gorm:"not null;comment:图书名称"`
	CoverURL    string `json:"coverURL" gorm:"comment:图书封面图片"`
	IsHidden    bool   `json:"isHidden" gorm:"not null;comment:是否显示;default: false"`
	Labels      string `json:"labels" gorm:"comment:图书标签"`
	Introduce   string `json:"introduce" gorm:"comment:图书简介;type:longtext"`
	SortID      uint   `json:"sort_id" gorm:"comment:分类id""`
	KeyWord     string `json:"keyWord" gorm:"comment:图书关键字"`
	Page        string `json:"page" gorm:"comment:图书页数"`
	Publisher   string `json:"publisher" gorm:"comment:图书出版社"`
	PublishTime string `json:"date" gorm:"comment:图书出版时间"`
	Author      string `json:"author" gorm:"comment:图书作者"`
}
