package domain

type UserID int

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type SessionID string
