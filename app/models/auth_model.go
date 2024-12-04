package models

type SignUpModel struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,password"`
	UserName string `json:"user_name" validate:"required"`
	Role     string `json:"role" validate:"required,userRole"`
}

type SignInModel struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"password,required"`
}
