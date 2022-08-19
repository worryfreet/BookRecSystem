package common

import "time"

// LibraryData 从超星数据库中爬取到的数据结构体
type LibraryData struct {
	Data struct {
		Result []struct {
			Author        string `json:"author"`
			BookCardD     string `json:"bookCardD"`
			BookName      string `json:"bookName"`
			CnFenlei      string `json:"cnFenlei"`
			CoverURL      string `json:"coverurl"`
			D             string `json:"d"`
			Dataproviders string `json:"dataproviders"`
			PublishTime   string `json:"date"`
			Dxid          string `json:"dxid"`
			Fenlei        string `json:"fenlei"`
			Introduce     string `json:"introduce"`
			IsFromBW      bool   `json:"isFromBW"`
			IsFromJG      bool   `json:"isFromJG"`
			IsFromNewLib  bool   `json:"isFromNewLib"`
			IsFromZX      bool   `json:"isFromZX"`
			JpathD        string `json:"jpathD"`
			KeyWord       string `json:"keyword"`
			Mulu          []struct {
				Dir     string `json:"dir"`
				PageNum int    `json:"pageNum"`
			} `json:"mulu"`
			Page       string `json:"page"`
			PdfOffline bool   `json:"pdfOffline"`
			PdgD       string `json:"pdgD"`
			Publisher  string `json:"publisher"`
			Ssid       string `json:"ssid"`
			ZxD        string `json:"zxD"`
		} `json:"result"`
		Total int `json:"total"`
	} `json:"data"`
	Success bool `json:"success"`
}

type Feedback struct {
	UserId   uint      `gorm:"comment:用户ID"`
	KeyWord  string    `gorm:"comment:图书ID"`
	SortId   uint      `gorm:"comment:分类"`
	FeedTime time.Time `gorm:"comment:反馈时间"`
	Score    int       `gorm:"comment:推荐程度"`
}
