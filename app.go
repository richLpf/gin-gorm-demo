package main

import (
	"fmt"
	"gin-gorm-demo/database"
	"gin-gorm-demo/models"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
)

type BuildConf struct {
	Mysql struct {
		Addr     string `yaml:"addr"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
	Listen string `yaml:"listen"`
	Env    string `yaml:"env"`
}

func (o *BuildConf) getConf() *BuildConf {
	yamlFile, err := ioutil.ReadFile("./build/dev.yaml")
	//yamlFile, err := ioutil.ReadFile("./build/prod.yaml")
	if err != nil {
		fmt.Println("err", err)
	}
	err = yaml.Unmarshal(yamlFile, o)
	if err != nil {
		fmt.Println("err", err)
	}
	return o
}

func main() {
	var err error

	var con BuildConf
	con.getConf()

	mysql_conf := con.Mysql
	connect_sql := mysql_conf.Username + ":" + mysql_conf.Password + "@tcp(" + mysql_conf.Addr + ":" + mysql_conf.Port + ")/" + mysql_conf.Database + "?"
	database.MYSQLDB, err = gorm.Open("mysql", connect_sql+"charset=utf8&parseTime=True&loc=Local")
	database.MYSQLDB.SingularTable(true) // User表表明默认为users,  如果设置了这一句，创建的表为user, 而不是用复数

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}

	if err != nil {
		fmt.Println("connection err")
	} else {
		switch con.Env {
		case "development":
			fmt.Println("current environment is development")
		case "production":
			fmt.Println("current environment is production")
		case "test":
			fmt.Println("current environment is test")
		}
	}
	defer database.MYSQLDB.Close()
	models.InitDB()

	router := InitRouter()
	fmt.Println("con", con.Listen)
	router.Run("0.0.0.0:" + con.Listen)
	fmt.Println("con", con.Listen)
}
