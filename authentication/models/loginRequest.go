package models

type LoginRequest struct {
	Email      string `json:"email" from:"email" binding:"required"`
	Password   string `json:"password" from:"password" binding:"required"`
	RememberMe bool   `json:"RememberMe" from:"RememberMe"`
}
