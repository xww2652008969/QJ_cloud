package global

import (
	"gitgub.com/xww2652008969/QJ-cloud/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var QjLog model.QjLog
var QjDb *gorm.DB
var QjRedis *redis.Client
