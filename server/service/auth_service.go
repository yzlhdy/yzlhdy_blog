package service

import (
	"fmt"
	"log"
	"yzlhdy_blog/dto"
	"yzlhdy_blog/entity"
	"yzlhdy_blog/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	// 登录
	VerifyCredential(email string, password string) interface{}
	// 注册
	CreateUser(user dto.UserCreateDto) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRep,
	}
}

func comparePassword(hash string, plainPassword []byte) bool {
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 登录
func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepo.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		compare := comparePassword(v.Password, []byte(password))
		if v.Email == email && compare {
			return res
		} else {
			return false
		}
	}
	return false
}

// 注册
func (service *authService) CreateUser(user dto.UserCreateDto) entity.User {
	userDtoCreate := entity.User{}
	err := smapping.FillStruct(&userDtoCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("failed to fill %v", err)
	}
	res := service.userRepo.InsertUser(userDtoCreate)
	return res
}
func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepo.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
