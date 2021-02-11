package users

// LoginRequest contains only fields a user needs to log in
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
