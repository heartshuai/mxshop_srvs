package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32     `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

/*
1.密文 2.不可反解

	1.对称加密
	2.非对称加密
	3.md5信息摘要算法

密码如果不可反解 用户找回密码 修改密码
*/
type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string     `gorm:"type:varchar(100);not null"`
	NickName string     `gorm:"type:varchar(20)"`
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;type:varchar(6);default:male;comment 'female表示女 male表示男'"`
	Role     int32      `gorm:"column:role;type:int;default:1;comment '1表示普通用户 2表示管理员'"`
}
