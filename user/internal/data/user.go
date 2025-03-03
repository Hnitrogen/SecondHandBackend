package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement;"`
	Name     string
	Email    string
	Password string
	Avatar   string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id uint) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id uint) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *User) error {
	return r.data.db.Create(user).Error
}

func (r *userRepo) GetUser(ctx context.Context, id uint) (*User, error) {
	var user User
	if err := r.data.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *User) error {
	return r.data.db.WithContext(ctx).Save(user).Error
}

func (r *userRepo) DeleteUser(ctx context.Context, id uint) error {
	return r.data.db.WithContext(ctx).Delete(&User{}, id).Error
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user *User
	if err := r.data.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
