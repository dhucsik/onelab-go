package models

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserRole  string `json:"user_role"`
}

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdatePasswordReq struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type UserBook struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Firsname string `json:"first_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Books    []Book `json:"books"`
}

type UserResp struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

const AdminRole = "admin"
const ModeratorRole = "moderator"
const UserRole = "user"
