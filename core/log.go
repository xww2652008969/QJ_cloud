package core

import (
	"gitgub.com/xww2652008969/QJ-cloud/global"
	"log"
)

func initlog() {
	global.QjLog.New("log.log", log.Lmicroseconds|log.Ldate)
	global.QjLog.SetLog()
	global.QjLog.Print("启动日志成功")
}
