package book_rec

import (
	"BookRecSystem/global"
	"BookRecSystem/model/book_rec"
	bookRecReq "BookRecSystem/model/book_rec/request"
	bookRecRes "BookRecSystem/model/book_rec/response"
	"BookRecSystem/model/common"
	"BookRecSystem/model/common/request"
	commonReq "BookRecSystem/model/common/request"
	"BookRecSystem/model/system"
	"BookRecSystem/utils"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type BookRecService struct {
}

func (b BookRecService) GetSysAllBooks(pageInfo request.PageInfo) (books []book_rec.Book, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	err = global.GSD_DB.Limit(limit).Offset(offset).Find(&books).Error
	return
}

func (b BookRecService) GetBookInfo(userId uint, bookFeed bookRecReq.GetBookWithId) (bookInfo book_rec.Book, err error) {
	if err = global.GSD_DB.Model(book_rec.Book{}).Where("id = ?", bookFeed.BookId).Find(&bookInfo).Error; err != nil {
		return
	}
	// 查询关键字是否存在, 存在则加分, 不存在则创建
	feedback := common.Feedback{
		UserId:   userId,
		KeyWord:  bookInfo.KeyWord,
		SortId:   bookInfo.SortID,
		FeedTime: time.Now(),
		Score:    0,
	}
	if err = global.GSD_DB.Model(&common.Feedback{}).Where("user_id = ? AND key_word = ?", userId, feedback.KeyWord).First(&feedback).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			global.GSD_DB.Model(&common.Feedback{}).Create(&feedback)
		} else {
			return
		}
	} else {
		global.GSD_DB.Model(&common.Feedback{}).Where("user_id = ? AND key_word = ?", userId, feedback.KeyWord).Update("score", feedback.Score+bookRecReq.FeedbackByDetailedInfo)
	}
	return bookInfo, nil
}

func (b BookRecService) UpdateBook(label bookRecReq.BookLabel) error {
	labelBytes, _ := json.Marshal(label.Labels)
	err := global.GSD_DB.Model(&book_rec.Book{}).Where("id = ?", label.BookId).Update("labels", string(labelBytes)).Error
	return err
}

func (b BookRecService) GetBookList(userId uint) (books []bookRecRes.BookShort, err error) {
	var booksId []uint
	err = global.GSD_DB.Model(&book_rec.BookList{}).Select("book_id").Where("user_id = ?", userId).Find(&booksId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if err = global.GSD_DB.Model(&book_rec.Book{}).
		Where("id in ?", booksId).Find(&books).Error; err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return books, nil
}

func (b BookRecService) GetBooksName(userId uint, query string) (booksName []bookRecRes.BookIdName, err error) {
	q := fmt.Sprint("%", query, "%")
	err = global.GSD_DB.Model(&book_rec.Book{}).Where("book_name like ? OR key_word like ?", q, q).
		Limit(30).Find(&booksName).Error
	// 查询关键字是否存在, 存在则加分, 不存在则创建
	var sortId uint
	err = global.GSD_DB.Model(&book_rec.Book{}).Select("sort_id").Where("book_name like ? OR key_word like ?", q, q).First(&sortId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	feedback := common.Feedback{
		UserId:   userId,
		KeyWord:  query,
		SortId:   sortId,
		FeedTime: time.Now(),
		Score:    0,
	}
	err = global.GSD_DB.Model(&common.Feedback{}).Where("user_id = ? AND key_word = ?", userId, feedback.KeyWord).First(&feedback).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			global.GSD_DB.Model(&common.Feedback{}).Create(&feedback)
		} else {
			return
		}
	} else {
		global.GSD_DB.Model(&common.Feedback{}).Where("user_id = ? AND key_word = ?", userId, feedback.KeyWord).Update("score", feedback.Score+bookRecReq.FeedbackByDetailedInfo)
	}
	return booksName, nil
}

func (b BookRecService) GetMyFavoriteBooks(userId uint) (books []bookRecRes.BookShort, err error) {
	if userId == 0 {
		global.GSD_DB.Model(&book_rec.Book{}).Limit(30).Find(&books)
		return
	}
	var labels []string
	var labelStr string
	global.GSD_DB.Model(system.SysUser{}).Where("id = ?", userId).Select("labels").Find(&labelStr)
	_ = json.Unmarshal([]byte(labelStr), &labels)
	// 根据图书关键字推荐
	for _, label := range labels {
		var fewBooks []bookRecRes.BookShort
		label = fmt.Sprint("%", label, "%")
		global.GSD_DB.Model(&book_rec.Book{}).
			Where("key_word like ? or book_name like ?", label, label).Limit(10).Find(&fewBooks)
		books = append(books, fewBooks...)
	}
	// 根据图书分类推荐
	queryMap := QueryMap()
	for i := 0; i < len(labels); i++ {
		var fewBooks []bookRecRes.BookShort
		global.GSD_DB.Model(&book_rec.Book{}).
			Where("sort_id = ?", queryMap[labels[i]]).Limit(10).Find(&fewBooks)
		books = append(books, fewBooks...)
	}
	// 根据反馈信息推荐
	var feedbacks []common.Feedback
	global.GSD_DB.Model(&common.Feedback{}).Where("user_id = ?", userId).
		Order("score desc").Limit(3).Find(&feedbacks)
	for i, fb := range feedbacks {

		// 根据反馈信息下的高频关键字进行推荐
		var fewBooksBySort []bookRecRes.BookShort
		global.GSD_DB.Model(&book_rec.Book{}).
			Where("sort_id = ?", queryMap[labels[i]]).Limit(10).Find(&fewBooksBySort)
		books = append(books, fewBooksBySort...)

		// 根据反馈信息下的高频图书分类进行推荐
		var fewBooksByKeyWord []bookRecRes.BookShort
		fb.KeyWord = fmt.Sprint("%", fb.KeyWord, "%")
		global.GSD_DB.Model(&book_rec.Book{}).
			Where("book_name like ? OR key_word like ?", fb.KeyWord, fb.KeyWord).Limit(10).Find(&fewBooksByKeyWord)
		books = append(books, fewBooksByKeyWord...)
	}
	return
}

func (b BookRecService) GetNewBooks(userId uint) (books []bookRecRes.BookShort, err error) {
	if userId == 0 {
		global.GSD_DB.Model(&book_rec.Book{}).Limit(30).Order("publish_time desc").Find(&books)
		return
	}
	var labels []string
	var labelStr string
	global.GSD_DB.Model(system.SysUser{}).Where("id = ?", userId).Select("labels").Find(&labelStr)
	_ = json.Unmarshal([]byte(labelStr), &labels)
	for i, label := range labels {
		var fewBooks []bookRecRes.BookShort
		label = fmt.Sprint("%", labels[i], "%")
		global.GSD_DB.Model(&book_rec.Book{}).
			Where("key_word like ? or book_name like ?", label, label).
			Order("publish_time desc").Limit(10).Find(&fewBooks)
		books = append(books, fewBooks...)
	}
	queryMap := QueryMap()
	for i := 0; i < len(labels); i++ {
		var fewBooks []bookRecRes.BookShort
		global.GSD_DB.Model(&book_rec.Book{}).
			Where("sort_id = ?", queryMap[labels[i]]).
			Order("publish_time desc").Limit(10).Find(&fewBooks)
		books = append(books, fewBooks...)
	}
	var rawBooks []bookRecRes.BookShort
	global.GSD_DB.Model(&book_rec.Book{}).Order("publish_time desc").Limit(10).Find(&rawBooks)
	books = append(books, rawBooks...)
	return
}

func (b BookRecService) AddBookIntoList(userId, bookId uint) (err error) {
	bookList := book_rec.BookList{
		UserID: userId,
		BookID: bookId,
	}
	err = global.GSD_DB.Create(&bookList).Error
	var ik bookRecReq.KeyAndSort
	if err = global.GSD_DB.Model(&book_rec.Book{}).Select("sort_id, key_word").Where("id = ?", bookId).First(&ik).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
	}
	feedback := common.Feedback{
		UserId:   userId,
		SortId:   ik.SortId,
		KeyWord:  ik.KeyWord,
		FeedTime: time.Now(),
		Score:    bookRecReq.FeedbackIntoBookList,
	}
	if err = global.GSD_DB.Model(&common.Feedback{}).Where("user_id = ? AND sort_id = ?", userId, ik.SortId).
		First(&feedback).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			global.GSD_DB.Model(&common.Feedback{}).Create(feedback)
		} else {
			return
		}
	} else {
		global.GSD_DB.Model(&common.Feedback{}).Where("user_id = ? AND sort_id = ?", userId, ik.SortId).
			UpdateColumn("score", feedback.Score+bookRecReq.FeedbackIntoBookList)
	}
	return nil
}

func (b BookRecService) DeleteBookFromList(userId uint, bookIds []uint) (err error) {

	err = global.GSD_DB.Model(&book_rec.BookList{}).Where("user_id = ? AND book_id in ?", userId, bookIds).Delete(&book_rec.BookList{}).Error
	var sortIds []uint
	if err = global.GSD_DB.Model(&book_rec.Book{}).Select("sort_id").Where("id in ?", bookIds).Find(&sortIds).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return
		}
	}
	var backs []common.Feedback
	global.GSD_DB.Model(&common.Feedback{}).Where("user_id = ? AND sort_id in ?", userId, sortIds).Find(&backs)
	for _, back := range backs {
		global.GSD_DB.Model(&common.Feedback{}).Where("user_id = ? AND sort_id in ?", userId, sortIds).
			Update("score", back.Score+bookRecReq.FeedbackOutBookList)
	}
	return nil
}

func (b BookRecService) GetBooksBySort(sortId bookRecReq.GetBooksWithSort) (books []bookRecRes.BookShort, err error) {
	global.GSD_DB.Model(&book_rec.Book{}).Where("sort_id = ?", sortId.SortId).Limit(40).Find(&books)
	return
}

func (b BookRecService) GetFeedBacks(pageInfo commonReq.PageInfo) (feedbacks []common.Feedback, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	err = global.GSD_DB.Limit(limit).Offset(offset).Find(&feedbacks).Error
	return
}

func QueryMap() map[string]int {
	ans := make(map[string]int, 22)
	for i, query := range utils.QueryList {
		ans[query] = i
	}
	return ans
}
