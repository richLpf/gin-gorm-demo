package common

import (
	"gin-gorm-demo/database"
)

/*判断表是否存在，如果不存在就创建*/
func HasTable(str string) bool {
	result := database.MYSQLDB.HasTable(str)
	return result
}
