package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 定义不同的模型

/*
* 定义公用模型
* 使用gorm来自定字段属性
 */
type Model struct {
	ID        uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id"` // primary_key定义主键   AUTO_INCREMENT定义自增
	CreatedAt time.Time  `json:"created_at,omitempty"`                 //omitempty 前端可以不用传入该字段，自动生成默认值，返回值如果为空，则不会返回该字段。
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

/*定义其他公共模型*/
type ResponseModel struct {
	Remark string `json:"remark"`
}

/*文章表结构*/
type Passages struct {
	gorm.Model         // 引用Model公共模型，使passages表中有ID CreatedAt  UpdatedAt  DeletedAt  字段， 不用重复写入
	Title       string `gorm:"not null;unique;" json:"title"` // 不为空，且唯一
	Author      string `json:"author"`
	Category    string `gorm:"size:255" json:"category"`            // 默认长度为255
	Tag         string `gorm:"type:varchar(100);unique" json:"tag"` //类型为varchar  最大100 列名为 `tag`
	LookNum     int    `json:"look,omitempty"`                      // 如果look_num的值为0  则不会返回该字段
	Description string `json:"description"`
	Content     string `json:"content,omitempty"` // 如果content字段为空或默认值，将不会返回改数据
	ImgLink     string `gorm:"-" json:"img_link"` // 忽略这个字段
}

/*注册用户表结构
* json: sex  和实际生成表列名没有关系，表列名字为 定义数据的蛇形小写
* 定义类型，type
* column 重新定义列的名字，如果像定义大写的列名字，经常用到
 */
type Users struct {
	gorm.Model
	Sex      int    // 列名为 `sex`  json仅仅为返回数据  可以不写默认返回大写字段可以不写
	Username string `json:"nick_name"` // 列名为 username
	Password string // 列名为password
	Remark   string `json:"description"`        // 列名为remark, json:"description" 为读取表数据后返回的字段写法
	Mail     string `gorm:"column:Mail"`        // 重新定义列名，Mail字段在表中的列名字 为Mail  大写的啊
	Phone    string `gorm:"column:MobilePhone"` // 该列的名字为 MobilePhone
	MoneyNum int    `gorm:"money_num"`          // 列名字为money_num
}

// 默认主键,更改表名
type LoginRecord struct {
	ID        uint // id 为默认主键
	CreatedAt time.Time
	Ip        string
}

// 返回数据结构引入提前定义的公共模型ResponseModel

type ResponseInfo struct {
	ResponseModel
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

/*定义表结构并重新命名, 将LoginRecord表名字重命名为 LoginLogs*/
func (LoginRecord) TableName() string {
	return "LoginLogs"
}
