package main

import (
	"BookRecSystem/core"
	"BookRecSystem/global"
	"BookRecSystem/initialize"
	"database/sql"
)

// @title                       BookRecSystem API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.GSD_VP = core.Viper()            // 初始化Viper
	global.GSD_LOG = core.Zap()             // 初始化zap日志库
	global.GSD_DB = initialize.Gorm()       // gorm连接数据库
	global.GSD_REDIS = initialize.Redis()   // 初始化redis
	global.GSD_Casbin = initialize.Casbin() // casbin初始化
	if global.GSD_DB != nil {
		//initialize.MysqlTables(global.GSD_DB) // 初始化表
		//defer global.GSD_WaitGroup.Wait()
		//utils.LibraryUpdate() // 更新数据库图书信息
		// 程序结束前关闭数据库链接
		db, _ := global.GSD_DB.DB()
		defer func(db *sql.DB) {
			_ = db.Close()
		}(db)
	}
	core.RunWindowsServer()
}
