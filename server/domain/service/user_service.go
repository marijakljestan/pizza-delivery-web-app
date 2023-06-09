package service

import (
	"errors"
	"fmt"
	"github.com/marijakljestan/golang-web-app/server/api/dto"
	domain "github.com/marijakljestan/golang-web-app/server/domain/model"
	repository "github.com/marijakljestan/golang-web-app/server/domain/repository"
	"github.com/marijakljestan/golang-web-app/server/helper/mapper"
	utils "github.com/marijakljestan/golang-web-app/server/util"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) RegisterCustomer(userDto dto.UserDto) (string, error) {
	if service.usernameExists(userDto.Username) {
		return "", errors.New(fmt.Sprintf("user with username %s already registered", userDto.Username))
	}
	user := mapper.MapUserToDomain(userDto)
	user.Role = domain.CUSTOMER
	user.Password = utils.HashPassword(user.Password)
	username, err := service.userRepository.Save(&user)
	if err != nil {
		fmt.Println(err)
	}
	return username, err
}

func (service *UserService) GetByUsername(username string) (*domain.User, error) {
	user, err := service.userRepository.GetByUsername(username)
	if err != nil {
		fmt.Println(err)
	}
	return user, err
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.userRepository.GetAll()
}

func (service *UserService) usernameExists(username string) bool {
	user, _ := service.userRepository.GetByUsername(username)
	if user == nil {
		return false
	}
	return true
}
