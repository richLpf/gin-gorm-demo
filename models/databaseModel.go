package models

import (
	"gin-gorm-demo/database"
)

/*定义模型*/
func AutoMigrateDB() {
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
