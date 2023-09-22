package schema

type SignupSchema struct {
	LoginSchema
	Name string `json:"name" validate:"required,min=3"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthResponse struct {
	User  *UserResponse `json:"user"`
	Token string        `json:"token"`
}

type LoginSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}
