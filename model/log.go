package model

import (
	"fmt"
	"gitgub.com/xww2652008969/QJ-cloud/utils"
	"log"
)

type QjLog struct {
	LogFileName string
	LogLogger   *log.Logger
	Logflag     int
}

//返回log文件地址
func (log1 *QjLog) Filename() string {
	return log1.LogFileName
}
func (log1 *QjLog) New(filename string, logflag int) {
	log1.LogFileName = filename
	log1.Logflag = logflag
}
func (log1 *QjLog) SetLog() {
	log1.LogLogger = utils.Setlog(log1.LogFileName, log1.Logflag, log1.LogLogger)
	fmt.Println(log1)
}
func (log1 *QjLog) Print(v ...any) {
	log1.LogLogger.Println(v)
}
