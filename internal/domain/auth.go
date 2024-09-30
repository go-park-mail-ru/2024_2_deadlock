package domain

type UserID int

type UserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,regexp=/^[a-zA-Z0-9?!_\\-*$]{6,255}$/"`
}

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type SessionID string
