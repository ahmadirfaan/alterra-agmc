package services

import (
	"alterra-agmc-dynamic-crud/models/database"
	"alterra-agmc-dynamic-crud/repositories"
)

type UserService interface {
	CreateNewUser(request database.User) error
	GetUserById(id int) (database.User, error)
	UpdateUser(user *database.User, id int) error
	DeleteUser(id int) error
	GetAllUsers(page int) ([]database.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(br repositories.UserRepository) UserService {
	return &userService{
		userRepository: br,
	}
}

func (b *userService) CreateNewUser(request database.User) error {
	_, err := b.userRepository.Save(request)
	return err
}

func (b *userService) GetUserById(id int) (database.User, error) {
	user, err := b.userRepository.FindByUserId(id)
	return user, err
}
func (b *userService) UpdateUser(user *database.User, id int) error {
	err := b.userRepository.UpdateUser(user, id)
	return err
}

func (b *userService) DeleteUser(id int) error {
	err := b.userRepository.DeleteUser(id)
	return err
}

func (b *userService) GetAllUsers(page int) ([]database.User, error) {
	var offset = 0
	if page > 1 {
		offset = 25 * page
	}
	users, err := b.userRepository.GetAllUsers(offset)
	return users, err
}
