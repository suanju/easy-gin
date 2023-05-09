package user

// RegisterReceive 用户注册
type RegisterReceive struct {
	UserName         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verificationCode" binding:"required"`
}

//LoginReceive 用户登入
type LoginReceive struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
