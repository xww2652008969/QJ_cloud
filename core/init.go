package core

func Init() {
	initlog() //初始化日式
	InitDB()  //初始化数据库
	//reids()      //初始化redis
	Migrate() //迁移数据库
}
