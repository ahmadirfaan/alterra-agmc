package services

import (
	"alterra-agmc-day3/models/database"
	models "alterra-agmc-day3/models/website"
	"alterra-agmc-day3/repositories"
	"alterra-agmc-day3/utils"
	"errors"
	"strconv"
)

type UserService interface {
	CreateNewUser(request models.CreateUserRequest) error
	GetUserById(id int) (database.User, error)
	UpdateUser(user *models.CreateUserRequest, idPath int, authorization string) error
	DeleteUser(idPath int, authorization string) error
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
func (b *userService) UpdateUser(user *models.CreateUserRequest, idPath int, authorization string) error {
	err4 := b.checkAuthenticationUser(idPath, authorization)
	if err4 != nil {
		return err4
	}
	newUser := &database.User{
		Name:     user.Name,
		Password: utils.HashPassword(user.Password),
		Email:    user.Email,
	}
	err := b.userRepository.UpdateUser(newUser, idPath)
	return err
}

func (b *userService) DeleteUser(idPath int, authorization string) error {
	err4 := b.checkAuthenticationUser(idPath, authorization)
	if err4 != nil {
		return err4
	}
	err := b.userRepository.DeleteUser(idPath)
	return err
}

func (b *userService) checkAuthenticationUser(idPath int, authorization string) error {
	userId, err3 := utils.ExtractToken(authorization)
	if err3 != nil {
		return err3
	}
	userIdInt, _ := strconv.Atoi(userId)
	userExist, err2 := b.GetUserById(idPath)
	if err2 != nil || (userExist == database.User{}) || userIdInt != *userExist.Id {
		return errors.New("403|NotAuthenticated")
	}
	return nil
}

func (b *userService) GetAllUsers(page int) ([]database.User, error) {
	var offset = 0
	if page > 1 {
		offset = 25 * page
	}
	users, err := b.userRepository.GetAllUsers(offset)
	return users, err
}

func (b *userService) UserLogin(username string) (string, error) {

	return username, nil
}
