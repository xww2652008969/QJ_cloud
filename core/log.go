package core

import (
	"gitgub.com/xww2652008969/QJ-cloud/global"
	"log"
)

func initlog() {
	global.QJ_log.New("log.log", log.Lmicroseconds|log.Ldate)
	global.QJ_log.SetLog()
}
