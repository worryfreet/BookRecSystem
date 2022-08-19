package book_rec

import (
	v1 "BookRecSystem/api"
	"github.com/gin-gonic/gin"
)

type BookRecRouter struct {
}

func (b *BookRecRouter) InitBookPubRecRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	bookRecRouter := Router.Group("recommend")
	var bookRecApi = v1.ApiGroupApp.BookRecApiGroup.BookRecApi
	{
		bookRecRouter.GET("bookInfo", bookRecApi.BookInfo)                   // 图书详情
		bookRecRouter.GET("booksNameBySearch", bookRecApi.BooksNameBySearch) // 查询获取图书名称列表
		bookRecRouter.GET("myFavoriteBooks", bookRecApi.MyFavoriteBooks)     // 我喜欢的书
		bookRecRouter.GET("newBooks", bookRecApi.NewBooks)                   // 最新推荐
		bookRecRouter.GET("booksBySort", bookRecApi.BooksBySort)
	}
	return bookRecRouter
}
func (b *BookRecRouter) InitBookRecPriRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	bookRecRouter := Router.Group("recommend")
	var bookRecApi = v1.ApiGroupApp.BookRecApiGroup.BookRecApi
	{
		bookRecRouter.GET("sysAllBooks", bookRecApi.SysAllBooks) // 后台获取全部图书
		bookRecRouter.PUT("updateBook", bookRecApi.UpdateBook)   // 编辑更新图书
		bookRecRouter.GET("bookList", bookRecApi.BookList)       // 书单
		bookRecRouter.POST("addBookIntoList", bookRecApi.AddBookIntoList)
		bookRecRouter.DELETE("deleteBookFromList", bookRecApi.DeleteBookFromList)
		bookRecRouter.GET("feedBacks", bookRecApi.Feedbacks)
	}
	return bookRecRouter
}
