package gorm

import (
	"easy-admin-go-service/models/system_models"
	"easy-admin-go-service/pkg/viper"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var BaseDataBase *gorm.DB

func init() {
	//	构造连接池链接
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.SrvConf.Database.Username,
		viper.SrvConf.Database.Password,
		viper.SrvConf.Database.Host,
		viper.SrvConf.Database.Db,
	)
	var err error
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	BaseDataBase, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	BaseDataBase.Logger.LogMode(logger.Info)
	if err != nil {
		log.Printf("The database connection failed may be due to a timeout")
		return
	}
	sqlDB, err := BaseDataBase.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Printf("The connection pool connection in the gorm database is successful")
	// 数据库迁移
	initDataBase()
	// 图标数据创建
	//InitSystemIcons()
}

// initDataBase 表迁移
func initDataBase() {
	var err error
	err = BaseDataBase.AutoMigrate(
		&system_models.SystemDepartmentModel{},
		&system_models.SystemMenuModel{},
		&system_models.SystemUserModel{},
		&system_models.SystemRoleModel{},
		&system_models.SystemRoleMenuModel{},
		&system_models.SystemRoleUserModel{},
	)

	if err != nil {
		log.Printf("initDataBase error %s", err.Error())
		return
	}
}
