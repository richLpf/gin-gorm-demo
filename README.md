## 项目简介

技术栈： go  gin  gorm

## api接口开发模板

- 发送邮件模块

- 文章列表增删改查功能

- 登录注册功能

## 之后完成的功能

- gorm各种操作方法应用

- gin路由配置，发布方法

- gin项目结构的搭建

目录结构,本项目结构，分成控制器，模型，数据库操作分离，这样项目体积增大后可以不至于太乱，变量容易污染。

之后优化方向：

- 一键打包，部署，用makefile管理
- 日志管理
- 数据库操作和api调用分离
- 公用函数分离

# 目录结构


- controller  控制器
    - admin.go   登录控制
    - passages.go  文章管理
    - region.go    机房管理
- database 数据库配置
- models  模型，处理数据库逻辑
- app.go  入口文件
- databaseModel.go  数据库结构体
- router.go  路由文件
- email.go  发送邮件服务

# 项目说明

> controller/region.go

api 透传，将要查找的数据，直接传递key,value  这样可以不需要判断，可以通过查询数据库，轻松的完成查询任务。





