package request

const (
	FeedbackBySort         = 1  // 通过查询进行反馈 +1
	FeedbackBySearch       = 3  // 通过查询进行反馈 +3
	FeedbackByDetailedInfo = 5  // 通过查看详情进行反馈 +5
	FeedbackIntoBookList   = 10 // 通过加入书单进行反馈 +10
	FeedbackOutBookList    = -3 // 通过从书单中删除进行反馈 -3
)

type BookLabel struct {
	BookId uint     `json:"book_id"`
	Labels []string `json:"labels"`
}

type GetBookWithId struct {
	BookId uint `json:"book_id" form:"book_id"`
}
type GetBookWithIds struct {
	BookIds []uint `json:"book_id" form:"book_id"`
}

type GetBooksWithSort struct {
	SortId uint `json:"sort_id" form:"sort_id"`
}

type KeyAndSort struct {
	KeyWord string `json:"key_word"`
	SortId  uint   `json:"sort_id"`
}

type Query struct {
	QueryWord string `json:"query" form:"query"`
}
