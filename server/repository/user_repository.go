package repository

import (
	"fmt"
	"log"
	"yzlhdy_blog/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUser(page int, limit int) []entity.User
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	DeleteUser(id int) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{connection: db}

}

// Password encryption
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to generate a password")
	}
	return string(hash)
}

func (db *userRepository) FindUser(page int, limit int) []entity.User {
	var users []entity.User
	db.connection.Limit(limit).Offset(page - 1).Find(&users)
	return users
}

func (db *userRepository) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Create(&user)
	return user
}

func (db *userRepository) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	}
	var userOld entity.User
	fmt.Println(user)
	db.connection.Model(userOld).Where("id = ?", user.ID).Updates(user)
	return userOld
}

// func (db *userRepository) FindUserById(id int) entity.User {
// 	var user entity.User
// 	db.connection.Where("id = ?", id).First(&user)
// 	return user
// }

func (db *userRepository) DeleteUser(id int) entity.User {
	var user entity.User
	db.connection.Delete(&user, id)
	return user
}

func (db *userRepository) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).First(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
func (db *userRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}
