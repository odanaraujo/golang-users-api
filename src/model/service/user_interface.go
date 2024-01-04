package service

import (
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/model"
	"github.com/odanaraujo/golang/users-api/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {

	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepo repository.UserRepository
}

type UserDomainService interface {
	CreateUser(user model.UserDomainInterface) (model.UserDomainInterface, *exception.Exception)
	UpdateUser(string, model.UserDomainInterface) *exception.Exception
	FindUserByID(id string) (model.UserDomainInterface, *exception.Exception)
	FindUserByEmail(id string) (*model.UserDomainInterface, *exception.Exception)
	DeleteUser(id string) *exception.Exception
	FindUserByID(string) (*model.UserDomainInterface, *exception.Exception)
	FindUserByEmail(string) (model.UserDomainInterface, *exception.Exception)
	DeleteUser(string) *exception.Exception
}
