package database

import "github.com/sho0126hiro/toriton/app/internal/interface/dao"

type UserRepository struct {
}

func (ur *UserRepository) Save(u dao.User) {}
