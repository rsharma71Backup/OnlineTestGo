package interfaces

import "OnlineTestGo/tos"

type AdminDao interface {
	Fetchdata() []tos.Admin
}
