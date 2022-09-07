package model

import "github.com/google/uuid"

type User struct {
	Id       uint `gorm:"primaryKey"`
	Uuid     uuid.UUID
	Pqssword string
	Name     string
	Email    string
}
