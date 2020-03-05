# Foxpot Backend by Go

## 项目结构
```
-foxpot_backend
    |-etc 系统配置
    |-models 应用数据库模型和其他共用结构体
    |-routers 路由逻辑处理
		|-api/v1 api
        |-middlewares 中间件
        |-views 页面路由
    |-statics 静态资源目录
        |-css css文件目录
        |-fonts 字体库
        |-images 图片目录
        |-js js文件目录
    |-utils 实用方法
    |-var 系统状态
		|-db 数据库
        |-log 日志
    |-views 模板文件
		|-admin 管理员页面
		|-root 根页面
		|-templates 模板文件
		|-user 普通用户页面
    |-main.go 程序执行入口
```

## TODO

- [x] 项目结构
- [x] 管理员和普通用户登录
- [x] 日志
- [x] 登录添加验证码
- [x] 配置文件
- [ ] 完善api 返回数据格式 错误码等等
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