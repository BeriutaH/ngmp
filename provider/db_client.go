package provider

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"ngmp/config"
	"ngmp/model"
	"strings"
	"time"
)

// Init 初始化 数据库链接
func Init() {
	// 读取数据库的配置文件
	InitConfig()
	config.DBDefault = gormMysql("default")
	config.RedisDefault = getRedisDb("default")

	// 数据库迁移
	err := config.DBDefault.AutoMigrate(&model.User{}, &model.Role{}, &model.Permission{}, &model.PermissionMenu{})
	fmt.Println("数据库迁移")
	if err != nil {
		fmt.Println(err)
	}

}

func getRedisDb(connection string) *redis.Client {
	connection = strings.ToUpper(connection)
	host := viper.GetString("Redis." + connection + ".Host")
	port := viper.GetInt("Redis." + connection + ".Port")
	database := viper.GetInt("Redis." + connection + ".Database")
	password := viper.GetString("Redis." + connection + ".Password")

	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Password:     password,
		DB:           database,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	if err := rdb.Ping(rdb.Context()).Err(); err != nil {
		fmt.Println("Redis start err: ", err, connection)
		return nil
	}
	return rdb
}

// MYSQL驱动
func gormMysql(connection string) *gorm.DB {
	connection = strings.ToUpper(connection)
	host := viper.GetString("DB." + connection + ".Host")
	port := viper.GetInt("DB." + connection + ".Port")
	database := viper.GetString("DB." + connection + ".Database")
	username := viper.GetString("DB." + connection + ".Username")
	password := viper.GetString("DB." + connection + ".Password")
	charset := viper.GetString("DB." + connection + ".Charset")
	log.Println("host", host)
	log.Println("port", port)
	log.Println("database", database)
	log.Println("username", username)
	log.Println("password", password)
	log.Println("charset", charset)

	// 拼接mysql相关配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)

	// 打开链接
	fmt.Println("mysql连接执行")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	sqlDB, _ := db.DB()
	// 设置数据库连接吃最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置连接池最大允许的空闲连接，如果没有sql执行，连接池多余20个的连接都被关闭
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

// InitConfig Config 配置
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic("Read config failed: " + err.Error())
	}

	// 监听配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: ", e.Name)
	})
}
