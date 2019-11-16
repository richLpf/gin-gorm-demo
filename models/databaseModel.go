package models

import (
	"gin-gorm-demo/database"
	"time"
)

/*登录记录*/
type LoginToken struct {
	Id        int       `json:"id,omitempty"`
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	ExpireAt  int64     `json:"expire_at"`
	Valid     int       `json:"valid,omitempty"`
}

/*定义模型*/
func InitDB() {
	// 自动迁移，会创建缺少的表，缺少的列和索引，不会改变现有的列类型或删除未使用的列
	database.MYSQLDB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").
		AutoMigrate(
			&Passages{},
			&Users{},
			&LoginToken{},
			&LoginRecord{},
			&ResponseInfo{},
		)
	database.MYSQLDB.DB().Ping()
}
