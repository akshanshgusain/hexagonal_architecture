package domain

// Server side PORT/ Secondary PORT

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}
