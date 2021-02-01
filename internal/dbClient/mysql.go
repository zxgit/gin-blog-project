package dbClient

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zxgit/gin-blog-project/global"
	"log"
	"time"
)

var Db *gorm.DB

func MysqlInit() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	dbType = global.DatabaseSetting.DBType
	dbName = global.DatabaseSetting.DBName
	user = global.DatabaseSetting.UserName
	password = global.DatabaseSetting.Password
	host = global.DatabaseSetting.Host
	tablePrefix = global.DatabaseSetting.TablePrefix

	Db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)

	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	//开启表明复数形式
	Db.SingularTable(true)
	//开启sql日志打印
	Db.LogMode(true)
	//最大空闲连接数
	Db.DB().SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	//最大打开连接数
	Db.DB().SetMaxOpenConns(global.DatabaseSetting.MaxOpenleConns)
	//创建更新前置时间戳
	Db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	Db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

}

func CloseDB() {
	defer Db.Close()
}

/**
 * @Author ZhangXin
 * @Description 前置回调更新时间/创建变更为时间戳
 * @Date 21:42 2021/1/24
 * @Param scope *gorm.Scope
 * @return
 **/
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createdAtField.IsBlank {
				fmt.Println(nowTime)
				createdAtField.Set(nowTime)
				fmt.Println(createdAtField.IsBlank)
			}
		}

		fmt.Println(333333)
		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(nowTime)
			}
		}
	}
}

/**
 * @Author ZhangXin
 * @Description  前置回调更新时间变更为时间戳
 * @Date 21:41 2021/1/24
 * @Param
 * @return
 **/
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	nowTime := time.Now().Unix()
	if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
		if updatedAtField.IsBlank {
			updatedAtField.Set(nowTime)
		}
	}
}
