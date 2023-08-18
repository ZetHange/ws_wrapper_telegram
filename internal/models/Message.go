package models

import "time"

type Update struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Author    Author    `json:"author"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type Author struct {
	ID           string    `json:"id"`
	Login        string    `json:"login"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	PdaID        int       `json:"pdaId"`
	Role         string    `json:"role"`
	Gang         string    `json:"gang"`
	XP           int       `json:"xp"`
	Registration time.Time `json:"registration"`
	LastLoginAt  time.Time `json:"lastLoginAt"`
}

type JSONData struct {
	Updates   []Update  `json:"updates"`
	Events    []Update  `json:"events"`
	Timestamp time.Time `json:"timestamp"`
}
