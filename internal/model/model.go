package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zxgit/gin-blog-project/global"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        int    `gorm:"primary_key" json:"id" form:"id"`
	CreatedAt uint32    `json:"createdAt"  form:"createdAt"`
	UpdatedAt uint32   `json:"updatedAt"  form:"updatedAt"`
	DeletedAt int    `json:"deletedAt"  form:"deletedAt"`
	IsDel     string `json:"is_del" form:"isDel"`

}

func (m *Model) GetDb() *gorm.DB {
	var db *gorm.DB
	return db
}

func (m *Model) DbInit() {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(sqlite3("test.db"), &gorm.Config{})
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

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
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

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(global.DatabaseSetting.MaxOpenleConns)
	db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
	//gorm.DefaultCallback.Create().Register("gorm:update_time_stamp", updateTimeStampForCreateCallback)

}

func CloseDB() {
	defer db.Close()
}
/**
 * @Author ZhangXin
 * @Description 前置回调更新时间/创建变更为时间戳
 * @Date 21:42 2021/1/24
 * @Param
 * @return
 **/
func updateTimeStampForCreateCallback(scope *gorm.Scope)  {
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
func updateTimeStampForUpdateCallback(scope *gorm.Scope)  {
	nowTime := time.Now().Unix()
	if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
		if updatedAtField.IsBlank {
			updatedAtField.Set(nowTime)
		}
	}
}
