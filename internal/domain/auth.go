package domain

type UserID int

type UserInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,regexp=/^[a-zA-Z0-9?!_\\-*$]{6,255}$/"`
}

type UserID int

type User struct {
	ID    UserID `json:"id"`
	Email string `json:"email"`
}

type SessionID string
