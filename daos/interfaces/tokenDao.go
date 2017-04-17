package interfaces

import (
	"OnlineTestGo/models"
	"time"
)

type TokenDao interface {
	Savetoken(token string, uid int64)
	ModifyToken(token string, uid int) error
	GetToken(token string, uid int64)
	AunthenticateToken(tokenEncodeString string) (string, time.Time)
	ModifyLastAccessTime(currentime time.Time, tokenEncodeString string) error
	DeleteToken(token models.Token, uid int64) error
}
