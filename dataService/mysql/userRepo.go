package mysql

import (
	"github.com/gjlee0802/my_introduce/domain/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo() *userRepo {
	return &userRepo{
		db: dbConn,
	}
}

func (ur *userRepo) GetByID(id int) (u *model.User, err error) {
	u = &model.User{}
	return u, ur.db.Where("id=?", id).First(u).Error
}

func (ur *userRepo) GetByName(name string) (u *model.User, err error) {
	u = &model.User{}
	return u, ur.db.Where("username=?", name).First(u).Error
}

func (ur *userRepo) Create(user *model.User) (*model.User, error) {
	return user, ur.db.Create(user).Error
}