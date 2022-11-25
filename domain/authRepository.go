package domain

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}
