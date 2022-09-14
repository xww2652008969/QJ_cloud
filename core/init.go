package core

import "gitgub.com/xww2652008969/QJ-cloud/utils"

func Init() {
	utils.Createfolder("work") //创建工作目录
	initlog()                  //初始化日式
	InitDB()                   //初始化数据库
	InitRedis()                //初始化redis
	Migrate()                  //迁移数据库
	Routers().Run(":8888")     //启动
}
