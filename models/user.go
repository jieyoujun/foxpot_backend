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
	Role           uint
	Email          string
	Phone          string
}

// HasAnyAdminByUsernames 是否包含管理员
func HasAnyAdminByUsernames(usernames []string) bool {
	for _, username := range usernames {
		if user, err := GetUserByUsername(username); err != nil && user.IsAdmin() {
			return true
		}
	}
	return false
}

// Role2Str 权限转文本
func Role2Str(role uint) string {
	switch {
	case role == 0:
		return "超级管理员"
	case role == 1:
		return "管理员"
	default:
		return "普通用户"
	}
}

// Role2Uint 权限转数字
func Role2Uint(role string) uint {
	switch {
	case role == "超级管理员":
		return uint(0)
	case role == "管理员":
		return uint(1)
	default:
		return uint(2)
	}
}

// IsAdmin 是否是管理员
func (u *User) IsAdmin() bool {
	return u.Role < 2
}

// CreateUser 新增用户
func CreateUser(user *User) error {
	return DB.Create(user).Error
}

// DeleteUser 删除用户
func DeleteUser(user *User) error {
	return DB.Delete(user).Error
}

// DeleteUserByUsername 通过用户名删除用户
func DeleteUserByUsername(userName string) error {
	return DB.Where("username = ?", userName).Delete(&User{}).Error
}

// UpdateUser 更新用户
func UpdateUser(user *User) error {
	return DB.Model(&User{}).Update(user).Error
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
