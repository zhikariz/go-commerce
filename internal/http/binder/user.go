package binder

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserCreateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
	NoHp     string `json:"no_hp" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UserUpdateRequest struct {
	ID       string `param:"id" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Alamat   string `json:"alamat" validate:"required"`
	NoHp     string `json:"no_hp" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UserDeleteRequest struct {
	ID string `param:"id" validate:"required"`
}

type UserFindByIDRequest struct {
	ID string `param:"id" validate:"required"`
}
