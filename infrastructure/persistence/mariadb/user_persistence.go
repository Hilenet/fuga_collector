package mariadb

import (
	"collector/domain/entity"
	"database/sql"
)

// respositoryのimpl
type UserPersistence struct {
	Conn *sql.DB
}

// conn受けて作る
func NewUserPersistence(conn *sql.DB) UserPersistence {
	return UserPersistence{conn}
}

func (p UserPersistence) Create(u *entity.User) (*entity.User, error) {
	return u, nil
}
func (p UserPersistence) Find(id int) (*entity.User, error) {
	return &entity.User{}, nil
}
func (p UserPersistence) FindbySlackID(slackID string) (*entity.User, error) {
	return &entity.User{}, nil
}
func (p UserPersistence) Update(u *entity.User) (*entity.User, error) {
	return u, nil
}
func (p UserPersistence) Delete(u *entity.User) (*entity.User, error) {
	return u, nil
}
