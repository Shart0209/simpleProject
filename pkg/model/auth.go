package model

type LoginRequest struct {
	Login    string `form:"username" binding:"required" json:"username"`
	Password string `form:"password" binding:"required" json:"password"`
}

type Auth struct {
	ID       string `json:"id,omitempty" db:"author_id"`
	Login    string `json:"login,omitempty" db:"login"`
	Name     string `json:"name,omitempty" db:"name"`
	Role     string `json:"role,omitempty" db:"role"`
	PswdHash string `json:"hash,omitempty" db:"pswd_hash"`
	Email    string `json:"email,omitempty" db:"email"`
}
