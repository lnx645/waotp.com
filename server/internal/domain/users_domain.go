package domain

type Users struct {
	Id         int    `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password"`
	Name       string `json:"name" db:"name"`
	IsVerified string `json:"is_verified" db:"is_verified"`
}
