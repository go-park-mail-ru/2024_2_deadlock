package domain

import "time"

type UserInputRegister struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,regexp=/^[a-zA-Z0-9?!_\\-*$]{6,255}$/"`
	FirstName string `json:"first-name" validate:"required, max=50"`
	LastName  string `json:"last-name" validate:"required, max=50"`
}

type UserInputLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,regexp=/^[a-zA-Z0-9?!_\\-*$]{6,255}$/"`
}

type UserID int

type User struct {
	ID        UserID  `json:"id"`
	Email     string  `json:"email"`
	AvatarURL *string `json:"avatar-url"`
	FirstName string  `json:"first-name"`
	LastName  string  `json:"last-name"`
}

type UserInfo struct {
	RegistrationDate time.Time `json:"registration-date"`
	ExtraInfo        *string   `json:"extra-info"`
	SubscribersNum   int       `json:"num-subscribers"`
	SubscriptionsNum int       `json:"num-subscriptions"`
	AvatarURL        *string   `json:"avatar-url"`
	FirstName        string    `json:"first-name"`
	LastName         string    `json:"last-name"`
}

type UserUpdate struct {
	Email            string  `json:"email" validate:"required,email"`
	ExtraInfo        *string `json:"extra-info" validate:"omitempty"`
	SubscribersNum   int     `json:"num-subscribers" validate:"required,gte=0"`
	SubscriptionsNum int     `json:"num-subscriptions" validate:"required,gte=0"`
	AvatarURL        *string `json:"avatar-url" validate:"required"`
	FirstName        string  `json:"first-name" validate:"required,max=50"`
	LastName         string  `json:"last-name" validate:"required,max=50"`
}

type PasswordUpdate struct {
	Password string `json:"password" validate:"required,regexp=/^[a-zA-Z0-9?!_\\-*$]{6,255}$/"`
}

type SessionID string
