package core

import (
	"BookRecSystem/global"
	"BookRecSystem/initialize"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GSD_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GSD_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 BookRecSystem
	当前版本:V1.0
	默认自动化文档地址:http://127.0.0.1%s/api/swagger/index.html
`, address)
	global.GSD_LOG.Error(s.ListenAndServe().Error())
}
