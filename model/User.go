package model

import "github.com/google/uuid"

type User struct {
	id       uint `gorm:"primaryKey"`
	Uuid     uuid.UUID
	PAssword string
	Name     string
	Email    string
}
