package model

import "github.com/google/uuid"

type User struct {
	Id         uint `gorm:"primaryKey"`
	Uuid       uuid.UUID
	Password   string
	Name       string
	Email      string
	Workaddres string
}
