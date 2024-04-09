package constants

type UserType string

const (
	UserContextKey = "user"
	AdminUser   UserType = "admin"
	RegularUser UserType = "user"
)
