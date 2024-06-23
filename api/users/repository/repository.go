package repository

import (
	"errors"

	"email.v1/api/users/dto"
	"email.v1/models"
	"gorm.io/gorm"
)

type repositoryEmail struct {
	db *gorm.DB
}

// FindUser implements RepositoryInterface.
type RepositoryInterface interface {
	Create(dto.RequestUser) error
	FindUserByEmail(email string) (dto.ResponseUser, error)
	Login(dto.RequestUser) (user models.Users, err error)
}

func NewRepositoryEmail(db *gorm.DB) RepositoryInterface {
	return &repositoryEmail{
		db: db,
	}
}

// register user
func (r *repositoryEmail) Create(input dto.RequestUser) error {
	dataInput := dto.RequestUserToModel(input)

	tx := r.db.Create(&dataInput)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// login user
func (r *repositoryEmail) Login(input dto.RequestUser) (user models.Users, err error) {
	// Cari pengguna berdasarkan email
	err = r.db.Where("username = ?", input.Password).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jika pengguna tidak ditemukan
			return user, errors.New("invalid email or password")
		}
		// Jika ada error lain dari database
		return user, err
	}
	// Kembalikan pengguna jika kredensial valid
	return user, nil
}

func (re *repositoryEmail) FindUserByEmail(email string) (dto.ResponseUser, error) {
	dataUser := models.Users{}

	tx := re.db.Where("email = ?", email).Take(&dataUser)
	if tx.Error != nil {
		return dto.ResponseUser{}, tx.Error
	}

	result := dto.ModelUserToResponse(dataUser)
	return result, nil
}
