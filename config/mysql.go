package config

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitConnect()  {
	dsn:="root:root@(127.0.0.1:3306)/test?charset=urtf8mb&parseTime=true&loc=Local"
	db,err:=gorm.Open("mysql",dsn)
	if err != nil {
		panic(err)
	}
	// 强制限制表明为自己定义的模型名单数形式
	db.SingularTable(true)
	_ = db.DB().Ping()
}

func CloseDB() {
	DB.Close()
}
func AutoMigrate() {
	DB.AutoMigrate()
}