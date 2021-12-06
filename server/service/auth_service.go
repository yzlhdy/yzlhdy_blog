package service

import (
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

func comparePassword(hashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Panic(err)
		return false
	}

	return true
}

// 登录
func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepo.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		compre := comparePassword(v.Password, []byte(password))
		if v, ok := res.(entity.User); ok && compre {
			return v
		}
		return res
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
