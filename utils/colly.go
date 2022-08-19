package utils

import (
	"BookRecSystem/global"
	"BookRecSystem/model/book_rec"
	"BookRecSystem/model/common"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var QueryList = []string{"马克思主义, 列宁主义", "哲学、宗教", "社会科学总论", "政治, 法律", "军事", "经济", "文化、科学、教育",
	"语言, 文字", "文学", "艺术", "历史、地理", "自然科学总论", "数理科学和化学", "天文学、地球科学", "生物科学", "医药、卫生", "农业科学",
	"工业技术", "交通、运输", "航空、航天", "环境科学、安全科学", "综合性文献"}

type Param struct {
	Sw           string `json:"sw"`
	Allsw        string `json:"allsw"`
	Searchtype   string `json:"searchtype"`
	ClassifyId   string `json:"classifyId"`
	Isort        string `json:"isort"`
	Field        uint   `json:"field"`
	Jsonp        string `json:"jsonp"`
	Showcata     string `json:"showcata"`
	Expertsw     string `json:"expertsw"`
	BCon         string `json:"BCon"`
	Page         string `json:"page"`
	Pagesize     uint   `json:"pagesize"`
	Sign         string `json:"sign"`
	Enc          string `json:"enc"`
	SearchNewLib uint   `json:"searchNewLib"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetBooksFromHTTP(classifyId string, page string, UA string) (res common.LibraryData) {
	client := &http.Client{}
	param := Param{ClassifyId: classifyId, Page: page}
	var data = strings.NewReader(ReqFormat(param))
	req, err := http.NewRequest("POST", "https://bookworld.sslibrary.com/book/search/do", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "account=; loginType=ip; deptid=1493; route=0b3146030a3862596ab641a58ad31944; JSESSIONID=F18E805D1DF7C7F3CABC8A0020BC352D.dsk44_web; msign=188892358556593; username=218%2e29%2e60%2e105; enc=922e770ef84c63ea22e9256e9e923f47; DSSTASH_LOG=C%5f34%2dUN%5f1493%2dUS%5f%2d1%2dT%5f1653134689313; ruot=1653134705558")
	req.Header.Set("Origin", "https://bookworld.sslibrary.com")
	req.Header.Set("Referer", "https://bookworld.sslibrary.com/book/search?classifyId="+classifyId)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", UA)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyText = []byte(strings.Replace(string(bodyText), `\`, "", -1))
	bodyText = []byte(strings.Trim(string(bodyText), `"`))
	err = json.Unmarshal(bodyText, &res)
	if err != nil {
		fmt.Println("bodyText Json Unmarshal failed, err: ", err)
		fmt.Println("-----------------------------------------")
	}
	return
}

func ReqFormat(param Param) string {
	reqStr := strings.Builder{}
	reqStr.WriteString("page=")
	reqStr.WriteString(param.Page)
	reqStr.WriteString("&")
	reqStr.WriteString("classifyId=")
	reqStr.WriteString(param.ClassifyId)
	reqStr.WriteString("&")
	reqStr.WriteString("pagesize=50&sw=&allsw=&searchtype=&isort=&field=1&jsonp=&showcata=&expertsw=&bCon=&sign=&enc=&searchNewLib=0")
	return reqStr.String()
}

func LibraryUpdate() {
	for i := 11; i <= 22; i++ {
		query := strconv.Itoa(i)
		if i < 10 {
			query = "0" + query
		}
		// 找到图书总数
		num := 10
		global.GSD_WaitGroup.Add(num)
		// 每次只能读取50页, 所以取除数+1
		for j := 1; j <= num; j++ {
			time.Sleep(50 * time.Millisecond)
			go func(q, page string, ii int) {
				data := GetBooksFromHTTP(q, page, RandomString())
				// 解析
				var books []book_rec.Book
				resBytes, err := json.Marshal(data.Data.Result)
				if err != nil {
					global.GSD_LOG.Error("colly json解析失败", zap.Any("err", err))
					return
				}
				err = json.Unmarshal(resBytes, &books)
				if err != nil {
					global.GSD_LOG.Error("colly json反序列化失败", zap.Any("err", err))
					return
				}
				for k := 0; k < len(books); k++ {
					books[k].SortID = uint(ii)
				}
				// 存进数据库
				global.GSD_DB.Model(&book_rec.Book{}).CreateInBatches(books, len(books))
				global.GSD_WaitGroup.Done()
			}(query, strconv.Itoa(j), i)
		}
	}
}
