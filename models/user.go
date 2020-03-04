package models

import (
	"time"
)

// User 系统用户
type User struct {
	ID             uint `gorm:"primary_key"`
	CreatedAt      time.Time
	LastLoginAt    time.Time
	Username       string
	HashedPassword string
	Role           string
	AvatarURL      string
	Email          string
	Phone          string
}

// IsAdmin 是否是管理员
func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

// CreateUser 新增用户
func CreateUser(user *User) error {
	return DB.Create(user).Error
}

// DeleteUser 删除用户
func DeleteUser(user *User) error {
	return DB.Delete(user).Error
}

// UpdateUser 更新用户
func UpdateUser(user *User) error {
	return DB.Save(user).Error
}

// GetUserByUsername 通过用户名获取用户
func GetUserByUsername(userName string) (*User, error) {
	var user User
	err := DB.Where("username = ?", userName).First(&user).Error
	return &user, err
}

// GetUserByID 通过用户ID获取用户
func GetUserByID(id uint) (*User, error) {
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

// GetAllUsers 获取所有用户
func GetAllUsers() ([]*User, error) {
	var users []*User
	err := DB.Find(&users).Error
	return users, err
}
