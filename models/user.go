package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	Username       string
	Password       string
	HashedPassword string
	Role           string
	AvatarURL      string
	Email          string
	Phone          string
	LastLoginAt    *time.Time
}

func (u *User) IsAdmin() bool {
	return u.Role == "administrator"
}

// 增
func InsertUser(user *User) error {
	return DB.Create(user).Error
}

// 删
func DeleteUser(user *User) error {
	return DB.Delete(user).Error
}

// 改
func UpdateUser(user *User) error {
	return DB.Save(user).Error
}

// 查
// SelectUserByUsername ...
func SelectUserByUsername(userName string) (*User, error) {
	var user User
	err := DB.Where("username = ?", userName).First(&user).Error
	return &user, err
}

// SelectUserByID ...
func SelectUserByID(id uint) (*User, error) {
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

// SelectAllUsers ...
func SelectAllUsers() ([]*User, error) {
	var users []*User
	err := DB.Find(&users).Error
	return users, err
}
