package repo

type TokenRepo interface {
	CreateToken() (*model.Token, error)
}
