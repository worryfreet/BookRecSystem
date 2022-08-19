package book_rec

import (
	"BookRecSystem/global"
	bookRecReq "BookRecSystem/model/book_rec/request"
	"BookRecSystem/model/common/request"
	"BookRecSystem/model/common/response"
	"BookRecSystem/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BookRecApi struct {
}

// @Tags Recommond
// @Summary 分页获取图书
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/sysAllBooks [get]
func (b BookRecApi) SysAllBooks(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userList, err := bookRecService.GetSysAllBooks(pageInfo)
	if err != nil {
		global.GSD_LOG.Error("获取失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(userList, c)
}

// @Tags Recommond
// @Summary 获取图书信息
// @accept application/json
// @Produce application/json
// @Param data query bookRecReq.GetBookWithId true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/bookInfo [get]
func (b BookRecApi) BookInfo(c *gin.Context) {
	userId := utils.GetUserID(c)
	var bookWithInfo bookRecReq.GetBookWithId
	_ = c.ShouldBindQuery(&bookWithInfo)
	if err := utils.Verify(bookWithInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bookInfo, err := bookRecService.GetBookInfo(userId, bookWithInfo)
	if err != nil {
		global.GSD_LOG.Error("获取图书信息失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取图书信息失败", c)
		return
	}
	response.OkWithData(bookInfo, c)
}

// @Tags Recommond
// @Summary 获取分类图书列表
// @accept application/json
// @Produce application/json
// @Param data query bookRecReq.GetBooksWithSort true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/booksBySort [get]
func (b BookRecApi) BooksBySort(c *gin.Context) {
	var bookWithSort bookRecReq.GetBooksWithSort
	_ = c.ShouldBindQuery(&bookWithSort)
	if err := utils.Verify(bookWithSort, utils.SortIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bookInfo, err := bookRecService.GetBooksBySort(bookWithSort)
	if err != nil {
		global.GSD_LOG.Error("获取图书信息失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取图书信息失败", c)
		return
	}
	response.OkWithData(bookInfo, c)
}

// @Tags Recommond
// @Summary 更新图书信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bookRecReq.BookLabel true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/updateBook [put]
func (b BookRecApi) UpdateBook(c *gin.Context) {
	var bookLabel bookRecReq.BookLabel
	_ = c.ShouldBindJSON(&bookLabel)
	if err := utils.Verify(bookLabel, utils.BookLabelVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := bookRecService.UpdateBook(bookLabel)
	if err != nil {
		global.GSD_LOG.Error("更新图书信息失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("更新图书信息失败", c)
		return
	}
	response.Ok(c)

}

// @Tags Recommond
// @Summary 获取个人书单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/bookList [get]
func (b BookRecApi) BookList(c *gin.Context) {
	userId := utils.GetUserID(c)
	bookList, err := bookRecService.GetBookList(userId)
	if err != nil {
		global.GSD_LOG.Error("获取书单失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取书单失败", c)
		return
	}
	response.OkWithData(bookList, c)
}

// @Tags Recommond
// @Summary 根据查询信息查找相关图书名称
// @accept application/json
// @Produce application/json
// @Param data query bookRecReq.Query true "查询字段"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/booksNameBySearch [get]
func (b BookRecApi) BooksNameBySearch(c *gin.Context) {
	userId := utils.GetUserID(c)
	var query bookRecReq.Query
	_ = c.ShouldBindQuery(&query)
	booksName, err := bookRecService.GetBooksName(userId, query.QueryWord)
	if err != nil {
		global.GSD_LOG.Error("查询失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(booksName, c)
}

// @Tags Recommond
// @Summary 推荐我可能喜欢的书
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/myFavoriteBooks [get]
func (b BookRecApi) MyFavoriteBooks(c *gin.Context) {
	userId := utils.GetUserID(c)
	books, err := bookRecService.GetMyFavoriteBooks(userId)
	if err != nil {
		global.GSD_LOG.Error("获取书单失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取书单失败", c)
		return
	}
	response.OkWithData(books, c)
}

// @Tags Recommond
// @Summary 系统推荐的新书
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/newBooks [get]
func (b BookRecApi) NewBooks(c *gin.Context) {
	userId := utils.GetUserID(c)
	books, err := bookRecService.GetNewBooks(userId)
	if err != nil {
		global.GSD_LOG.Error("获取书单失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取书单失败", c)
		return
	}
	response.OkWithData(books, c)
}

// @Tags Recommond
// @Summary 添加图书到书单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bookRecReq.GetBookWithId true "图书id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/addBookIntoList [post]
func (b BookRecApi) AddBookIntoList(c *gin.Context) {
	userId := utils.GetUserID(c)
	var bookWithInfo bookRecReq.GetBookWithId
	_ = c.ShouldBindJSON(&bookWithInfo)
	if err := utils.Verify(bookWithInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := bookRecService.AddBookIntoList(userId, bookWithInfo.BookId)
	if err != nil {
		global.GSD_LOG.Error("添加书单失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("添加书单失败", c)
		return
	}
	response.Ok(c)
}

// @Tags Recommond
// @Summary 从书单中删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bookRecReq.GetBookWithIds true "图书id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/deleteBookFromList [delete]
func (b BookRecApi) DeleteBookFromList(c *gin.Context) {
	userId := utils.GetUserID(c)
	var bookWithIds bookRecReq.GetBookWithIds
	_ = c.ShouldBindJSON(&bookWithIds)
	err := bookRecService.DeleteBookFromList(userId, bookWithIds.BookIds)
	if err != nil {
		global.GSD_LOG.Error("从书单中删除失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("从书单中删除失败", c)
		return
	}
	response.Ok(c)
}

// @Tags Recommond
// @Summary 后台获取反馈信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/recommend/feedBacks [get]
func (b BookRecApi) Feedbacks(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	feedbacks, err := bookRecService.GetFeedBacks(pageInfo)
	if err != nil {
		global.GSD_LOG.Error("获取反馈信息失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取反馈信息失败", c)
		return
	}
	response.OkWithData(feedbacks, c)
}
