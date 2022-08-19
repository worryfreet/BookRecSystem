package initialize

import (
	"BookRecSystem/global"
	"BookRecSystem/model/book_rec"
	"BookRecSystem/model/common"
	"BookRecSystem/model/system"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// Gorm
// @description: 初始化数据库并产生数据库全局变量
// @return: *gorm.DB
func Gorm() *gorm.DB {
	return GormMysql()
}

// MysqlTables
// @description: 注册数据库表专用
// @param: db *gorm.DB
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysApi{},
		system.CasbinModel{},
		system.SysAuthority{},
		book_rec.Book{},
		book_rec.BookList{},
		common.Feedback{},
	)
	if err != nil {
		global.GSD_LOG.Error("Register Table Failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GSD_LOG.Info("Register Table Success")
}

// GormMysql
// @description: 初始化Mysql数据库
// @return: *gorm.DB
func GormMysql() *gorm.DB {
	m := global.GSD_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		global.GSD_LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config
func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GSD_CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = Default.LogMode(logger.Info)
	default:
		config.Logger = Default.LogMode(logger.Info)
	}
	return config
}
