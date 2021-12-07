package service

import (
	"log"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/entity"
	"yzlhdy_blog/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	FindUser(page int, limit int) []entity.User
	UpdateUser(user dto.UserUpdateDto) entity.User
	DeleteUser(id int) entity.User
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) FindUser(page int, limit int) []entity.User {
	return u.userRepo.FindUser(page, limit)
}

func (u *userService) UpdateUser(user dto.UserUpdateDto) entity.User {
	userData := entity.User{}
	err := smapping.FillStruct(&userData, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updateUser := u.userRepo.UpdateUser(userData)
	return updateUser
}

func (u *userService) DeleteUser(id int) entity.User {
	return u.userRepo.DeleteUser(id)
}
