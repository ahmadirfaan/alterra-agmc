package services

import (
	"alterra-agmc-day3/models/database"
	models "alterra-agmc-day3/models/website"
	"alterra-agmc-day3/repositories"
)

type UserService interface {
	CreateNewUser(request models.CreateUserRequest) error
	GetUserById(id int) (database.User, error)
	UpdateUser(user *models.CreateUserRequest, id int) error
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

func (b *userService) CreateNewUser(request models.CreateUserRequest) error {
	newUser := database.User{
		Name:     request.Name,
		Password: request.Password,
		Email:    request.Email,
	}
	_, err := b.userRepository.Save(newUser)
	return err
}

func (b *userService) GetUserById(id int) (database.User, error) {
	user, err := b.userRepository.FindByUserId(id)
	return user, err
}
func (b *userService) UpdateUser(user *models.CreateUserRequest, id int) error {
	newUser := &database.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}
	err := b.userRepository.UpdateUser(newUser, id)
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
