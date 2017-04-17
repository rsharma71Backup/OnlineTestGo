package interfaces

type LogoutDao interface {
	DeleteToken(token string) string
}
