// internal/biz/user.go
package biz

import (
	"context"
	"errors"
	"time"
	"user/internal/data"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt"
)

type UserUseCase struct {
	repo data.UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo data.UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *data.User) error {
	// find user by email
	db_user, err := uc.repo.GetUserByEmail(ctx, user.Email)

	// check if user already exists
	if db_user != nil || err == nil {
		return errors.New("用户已存在")
	}

	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id uint) (*data.User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, user *data.User) error {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id uint) error {
	return uc.repo.DeleteUser(ctx, id)
}

func (uc *UserUseCase) UserLogin(ctx context.Context, email string, password string) (map[string]interface{}, error) {
	// find user by email
	user, err := uc.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	// check password
	if user.Password != password {
		return nil, errors.New("password is incorrect")
	} else {
		// Jwt 生成token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":    email,
			"username": user.Name,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"token": tokenString,
			"user":  user,
		}, nil
	}
}
