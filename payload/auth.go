package payload

type RegisterUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Nik      string `json:"nik" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Nik      string `json:"nik"`
	Password string `json:"password" binding:"required"`
}
