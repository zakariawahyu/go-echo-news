package entity

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64  `bun:",pk" json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
}

type UserResponse struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64  `bun:",pk" json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
}
