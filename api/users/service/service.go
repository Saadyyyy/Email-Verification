package service

import (
	"errors"
	"fmt"

	"email.v1/api/users/dto"
	"email.v1/api/users/repository"
	"email.v1/models"
	"email.v1/utils/helper"
	"email.v1/utils/validation"
	"gorm.io/gorm"
)

type ServiceEmail struct {
	db         *gorm.DB
	repository repository.RepositoryInterface
}

// FindUser implements RepositoryInterface.
type ServiceInterface interface {
	Create(input dto.RequestUser) error
	FindUserByEmail(email string) (dto.ResponseUser, error)
	Login(input dto.RequestUser) error
}

func NewRepositoryEmail(repository repository.RepositoryInterface, db *gorm.DB) ServiceInterface {
	return &ServiceEmail{
		repository: repository,
		db:         db,
	}
}

func (s *ServiceEmail) Create(input dto.RequestUser) error {

	errEmpty := validation.CheckEmpty(input.Email, input.Password, input.Username)
	if errEmpty != nil {
		return errEmpty
	}

	_, err := s.repository.FindUserByEmail(input.Email)
	if err == nil {
		return errors.New("errors : email already exist")
	}

	if input.Username != "" {
		err := s.db.Where("username = ?", input.Username).First(&models.Users{}).Error
		if err == nil {
			return errors.New("error: username already exists")
		}
	}

	Encrypt, err := helper.HashPassword(input.Password)
	if err != nil {
		return errors.New("errors : failed to encrypt password")
	}

	input.Password = Encrypt
	user := models.Users{}

	if !user.IsVerify {
		return fmt.Errorf("Harap register lebih dahulu")
	}

	err = s.repository.Create(input)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceEmail) Login(input dto.RequestUser) error {
	return nil
}

func (s *ServiceEmail) FindUserByEmail(email string) (dto.ResponseUser, error) {

	result, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return dto.ResponseUser{}, err
	}
	return result, nil

}
