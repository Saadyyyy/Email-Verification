package service

import (
	"errors"

	"email.v1/api/dto"
	"email.v1/api/repository"
	"email.v1/utils/helper"
	"email.v1/utils/validation"
)

type ServiceEmail struct {
	repository repository.RepositoryInterface
}

// FindUser implements RepositoryInterface.
type ServiceInterface interface {
	CreateUser(input dto.RequestUser) error
	FindUserByEmail(email string) (dto.ResponseUser, error)
}

func NewRepositoryEmail(repository repository.RepositoryInterface) ServiceInterface {
	return &ServiceEmail{
		repository: repository,
	}
}

func (s *ServiceEmail) CreateUser(input dto.RequestUser) error {

	errEmpty := validation.CheckEmpty(input.Email, input.Password, input.Username)
	if errEmpty != nil {
		return errEmpty
	}

	_, err := s.repository.FindUserByEmail(input.Email)
	if err == nil {
		return errors.New("errors : email already exist")
	}

	Encrypt, err := helper.HashPassword(input.Password)
	if err != nil {
		return errors.New("errors : failed to encrypt password")
	}

	input.Password = Encrypt

	err = s.repository.CreateUser(input)
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceEmail) FindUserByEmail(email string) (dto.ResponseUser, error) {

	result, err := s.repository.FindUserByEmail(email)
	if err != nil {
		return dto.ResponseUser{}, err
	}
	return result, nil

}
