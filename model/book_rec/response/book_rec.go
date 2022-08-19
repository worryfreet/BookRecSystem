package response

type BookIdName struct {
	BookName string `json:"book_name"`
	Id       uint   `json:"book_id"`
}

type BookShort struct {
	BookIdName
	CoverUrl  string `json:"cover_url"`
	Author    string `json:"author"`
	Introduce string `json:"introduce"`
}
