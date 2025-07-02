package repository

type IQuery interface {
	CreateUser() string
	SaveRefreshToken() string
}
