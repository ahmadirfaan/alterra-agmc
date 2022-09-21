package repositories

import (
	"alterra-agmc-day6/models/database"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user database.User) (database.User, error)
	FindByUserId(userId int) (database.User, error)
	UpdateUser(user *database.User, id int) error
	DeleteUser(id int) error
	GetAllUsers(page int) ([]database.User, error)
	UserLogin(email string) (database.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u userRepository) Save(user database.User) (database.User, error) {
	err := u.DB.Debug().Save(&user).Error
	fmt.Println(err)
	log.Printf("Users Repositories:%+v\n", user)
	return user, err
}

func (u userRepository) UpdateUser(user *database.User, id int) error {
	if err := u.DB.Model(user).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (u userRepository) FindByUserId(userId int) (database.User, error) {
	var user database.User
	err := u.DB.Where("id = ?", userId).First(&user).Error
	return user, err
}

func (u userRepository) DeleteUser(id int) error {
	if err := u.DB.Delete(&database.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (u userRepository) GetAllUsers(page int) ([]database.User, error) {
	var users []database.User
	var offset = 0
	if page > 1 {
		offset = 25 * page
	}

	result := u.DB.Limit(25).Offset(offset).Find(&users)

	return users, result.Error
}

func (u userRepository) UserLogin(email string) (database.User, error) {
	var user database.User
	result := u.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}
