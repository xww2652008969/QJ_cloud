package service

type GormService struct {
	UserGormapi UserGormService
}
type RedisSevice struct {
}
type Service struct {
	Gorm  GormService
	Redis RedisSevice
}

var ServiceGroup = new(Service)
