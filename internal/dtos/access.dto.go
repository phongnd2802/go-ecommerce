package dtos

import "time"

type (
	ShopRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	TokenReponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	ShopResponse struct {
		ID        string       `json:"id"`
		Name      string       `json:"name"`
		Email     string       `json:"email"`
		IsActive  bool         `json:"is_active"`
		CreatedAt time.Time    `json:"created_at"`
		UpdatedAt time.Time    `json:"updated_at"`
		Tokens    TokenReponse `json:"tokens"`
	}

)
