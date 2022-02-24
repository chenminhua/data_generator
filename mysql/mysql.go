package mysql

import (
	"database/sql"
	"github.com/chenminhua/data_generator/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func OpenDB() {
	sqlDB, err := sql.Open("mysql", "root:secretpassword@tcp(127.0.0.1:3306)/test?parseTime=true")
	DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Silent),
		CreateBatchSize: 100,
	})
	println(DB.CreateBatchSize)
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&types.User{})
}

func CreateUser(start, end int) {
	users := []*types.User{}
	for idx := start; idx < end; idx++ {
		users = append(users, types.NewUser(idx))
	}
	DB.Create(users)
}
