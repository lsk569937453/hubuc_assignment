package vojo

type RegisterReq struct {
	UserName *string `form:"userName" binding:"required"`
	Email    *string `form:"email" binding:"required,email"  json:"userOldPassword"`
	Password *string `form:"password" binding:"required"  json:"userNewPassword"`
}
