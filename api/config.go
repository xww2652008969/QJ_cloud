package api

import v1 "gitgub.com/xww2652008969/QJ-cloud/api/v1"

type ApiGroup struct {
	Userapi v1.UserApi
}

var ApigroupApp = new(ApiGroup)
