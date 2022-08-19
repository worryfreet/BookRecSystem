package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// OperateRequestLog
// @description: 处理请求数据
// @param: c *gin.Context
// @return: mBody []byte post方法参数, queryGet []byte get方法参数
func OperateRequestLog(c *gin.Context) ([]byte, []byte) {
	var mBody []byte
	var query string
	var queryGet []byte
	if c.Request.Method != http.MethodGet {
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		var mPost map[string]string
		_ = json.Unmarshal(body, &mPost)
		for k := range mPost {
			if k == "phone" {
				delete(mPost, k)
			}
		}
		mBody, _ = json.Marshal(mPost)
	} else {
		query = c.Request.URL.RawQuery
		query, _ = url.QueryUnescape(query)
		split := strings.Split(query, "&")
		mGet := make(map[string]string)
		for _, v := range split {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				mGet[kv[0]] = kv[1]
			}
		}
		queryGet, _ = json.Marshal(mGet)
	}
	return mBody, queryGet
}
