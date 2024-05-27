package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	Id        string
	Username  string
	Password  string
	IsVerify  bool
	CreatedAt time.Time
}

type TokenVerif struct {
	Id        string
	UsersID   string
	Token     string
	Expire    int64
	CreatedAt time.Time
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	u.Id = newUuid.String()

	return nil
}

func (t *TokenVerif) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	t.Id = newUuid.String()

	return nil
}
