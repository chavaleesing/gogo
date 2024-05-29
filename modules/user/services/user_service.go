package services

import (
	"fmt"

	"gogo/modules/entities"

	"golang.org/x/crypto/bcrypt"
)

type userUse struct {
	UsersRepo entities.UserRepository
}

// Constructor
func NewUserService(usersRepo entities.UserRepository) entities.UserService {
	return &userUse{
		UsersRepo: usersRepo,
	}
}

func (u *userUse) Register(req *entities.UserRegisterReq) (*entities.UserRegisterRes, error) {
	// Hash a password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	req.Password = string(hashed)

	// Send req next to repository
	user, err := u.UsersRepo.Register(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
