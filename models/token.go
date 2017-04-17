package models

import (
	"time"
)

type Token struct {
	ID             int       `json:"id"`
	Uid            int       `json:"uid"`
	Token          string    `json:"token"`
	LastAccessTime time.Time `json:"lastaccesstime"`
}
