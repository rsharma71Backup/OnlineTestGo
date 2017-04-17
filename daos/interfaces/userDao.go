package interfaces

import (
	"OnlineTestGo/models"
	"OnlineTestGo/tos"
)

type UserDao interface {
	SaveNewUser(u models.User) (tos.Userto, error)
	CheckUser(u models.User) (tos.Userto, error)
	AuthenticateUser(user models.User) models.User
}
