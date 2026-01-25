package dto

type (
	RegisterRequest struct {
		Username        string `json:"username" validate:"required"`
		Password        string `json:"password" validate:"required"`
		PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	}

	RegisterResponse struct {
		UserID int64 `json:"user_id"`
	}

	LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Token    string `json:"token"`
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`
		//RefreshToken string `json:"refresh_token"`
	}
)
