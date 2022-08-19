package middleware

import (
	"BookRecSystem/global"
	"BookRecSystem/utils"
	"bytes"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	// b 就是 response
	// Write 方法把 response 缓存到 responseBodyWriter 结构体的 body 属性中
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinBodyLogMiddleware(c *gin.Context) string {
	// 使用 responseBodyWriter 替换 gin 中的 responseWriter,
	// 替换的目的是把 response 返回值缓存起来
	w := &bodyLogWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w
	c.Next()
	return w.body.String()
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		requestId, _ := gonanoid.Nanoid()
		c.Set("requestId", requestId)
		mBody, queryGet := utils.OperateRequestLog(c)
		cost := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		ip := c.ClientIP()
		errString := c.Errors.ByType(gin.ErrorTypePrivate).String()
		userAgent := c.Request.UserAgent()
		global.GSD_LOG.Info(path,
			zap.String("requestId", requestId),
			zap.Int("status", status),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("query", string(queryGet)),
			zap.String("body", string(mBody)),
			zap.String("ip", ip),
			zap.String("user-agent", userAgent),
			zap.String("errors", errString),
			zap.Duration("cost", cost),
		)
		type response struct {
			Code int         `json:"code"`
			Data interface{} `json:"data"`
			Msg  string      `json:"msg"`
		}
		resp := GinBodyLogMiddleware(c)
		var respStruct response
		_ = json.Unmarshal([]byte(resp), &respStruct)
		global.GSD_LOG.Info(path,
			zap.Int("code", respStruct.Code),
			zap.String("msg", respStruct.Msg),
			zap.Any("data", respStruct.Data),
		)
	}
}
