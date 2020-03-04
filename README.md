# Foxpot Backend by Go

## TODO

- [x] 项目结构
- [x] 管理员和普通用户登录
- [x] 日志
- [ ] 登录添加验证码
- [ ] 配置自动
- [ ] 完善api 返回数据格式 错误码等等
- [ ] 添加redis管理session
- [ ] ... 

## Use

- Gin
- Gin sessions
- gorm

## 插入测试用户
```
hashedPassword, _ := utils.HashPassword("1212")
models.CreateUser(&models.User{
	Username:       "liki",
	HashedPassword: hashedPassword,
	Role:           "admin",
	Email:          "admin@foxpot.com",
	Phone:          "01234567890",
})
models.CreateUser(&models.User{
	Username:       "niki",
	HashedPassword: hashedPassword,
	Role:           "user",
	Email:          "user@foxpot.com",
	Phone:          "01234567890",
})
```