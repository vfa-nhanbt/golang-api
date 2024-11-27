package models

type SignUpModel struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}

type SignInModel struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"password,required"`
}
