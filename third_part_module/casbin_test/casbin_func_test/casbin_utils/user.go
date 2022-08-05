package casbin_utils

type User interface {
	GetID() int
	GetUsername() string
	GetGroup() string
	GetRole() string
}
