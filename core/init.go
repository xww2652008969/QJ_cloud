package core

import "gitgub.com/xww2652008969/QJ-cloud/global"

func Init() {
	initlog()
	global.QJ_log.Print("日志系统成功启动")
}
