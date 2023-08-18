package models

import "time"

type User struct {
	ID           string    `json:"id"`
	Login        string    `json:"login"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	PdaID        int       `json:"pdaId"`
	Role         string    `json:"role"`
	Gang         string    `json:"gang"`
	Xp           int       `json:"xp"`
	Registration time.Time `json:"registration"`
	LastLoginAt  time.Time `json:"lastLoginAt"`
}
