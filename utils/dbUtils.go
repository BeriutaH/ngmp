package utils

import (
	"fmt"
	"time"
)

/*此文件redis与mysql配置相关*/

var (
	db    *gorm.DB
	Redis *redis.Client
)

// Init 初始化 数据库链接
func Init() {
	var err error

	// mysql 数据库链接 database, err := gorm.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	fmt.Println("mysql连接执行")
	constr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456789", "127.0.0.1", 3306, "ngmp")
	db, err = gorm.Open(mysql.Open(constr), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	// 数据库迁移
	err = db.AutoMigrate(&User{}, &Role{}, &Permission{}, &PermissionMenu{})
	fmt.Println("数据库迁移")
	if err != nil {
		fmt.Println(err)
	}

	// 设置数据库连接池参数
	sqlDB, _ := db.DB()
	// 设置数据库连接吃最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置连接池最大允许的空闲连接，如果没有sql执行，连接池多余20个的连接都被关闭
	sqlDB.SetMaxIdleConns(20)

	// redis数据库链接
	Redis = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

func GetDB() *gorm.DB {
	return db
}
