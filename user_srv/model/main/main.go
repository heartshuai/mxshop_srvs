package main

import (
	"crypto/md5"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"mxshop_srvs/user_srv/model"
	"os"
	"time"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}
func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "mxshop_user_srv:xnihiwhd8EbtFsMB@tcp(152.136.246.107:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "mxshop_",
			SingularTable: true,
			//NameReplacer:  nil,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	///**
	//设置全局的logger 这个logger在我们执行每个sql语句的时候会打印每一行sql
	//sql才是最重要的，本着这个原则尽量看到每个api背后的sql语句是什么样的
	//定义一个表结构 将表结构直接生成对应的表 -migrations
	//*/
	//// 迁移 schema
	//db.AutoMigrate(&model.User{}) //此处应该有sql语句
	//salt, encodedPwd := password.Encode("generic password", nil)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(salt, encodedPwd, check) // true

	// Using custom options
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("admin123", options)
	newPassword := fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodedPwd)
	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("yxs%d", i),
			Mobile:   fmt.Sprintf("1840165443%d", i),
			Password: newPassword,
		}
		db.Save(&user)

	}

	//fmt.Println(len(newPassword))
	//passwordInfo := strings.Split(newPassword, "$")
	//salt = passwordInfo[2]
	//encodedPwd = passwordInfo[3]
	//check := password.Verify("generic password", salt, encodedPwd, options)
	////fmt.Println(salt, encodedPwd, check) // true
	//fmt.Println(check)

}
