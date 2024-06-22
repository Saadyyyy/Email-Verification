package repository

import (
	"email.v1/api/dto"
	"email.v1/models"
	"gorm.io/gorm"
)

type repositoryEmail struct {
	repository *gorm.DB
}

// FindUser implements RepositoryInterface.
type RepositoryInterface interface {
	CreateUser(dto.RequestUser) error
	FindUserByEmail(email string) (dto.ResponseUser, error)
}

func NewRepositoryEmail(repository *gorm.DB) RepositoryInterface {
	return &repositoryEmail{
		repository: repository,
	}
}

func (re *repositoryEmail) CreateUser(input dto.RequestUser) error {
	dataInput := dto.RequestUserToModel(input)

	tx := re.repository.Create(&dataInput)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (re *repositoryEmail) FindUserByEmail(email string) (dto.ResponseUser, error) {
	dataUser := models.Users{}

	tx := re.repository.Where("email = ?", email).Take(&dataUser)
	if tx.Error != nil {
		return dto.ResponseUser{}, tx.Error
	}

	result := dto.ModelUserToResponse(dataUser)
	return result, nil
}
