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