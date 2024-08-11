package dtos

type (
	ShopRegisterRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	ShopResponse struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
