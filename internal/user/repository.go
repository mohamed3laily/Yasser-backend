package user

import (
	"yasser-backend/database"
)

type Repository interface {
	FindByPhone(phone string) (*User, error)
	Create(phone string) (*User, error)
	GetByID(id uint) (*User, error)
	Update(user *User) (*User, error)
	Delete(id uint) error
}

type repo struct{}

func NewRepository() Repository {
	return &repo{}
}

func (r *repo) FindByPhone(phone string) (*User, error) {
	var user User
	result := database.DB.Where("phone_number = ?", phone).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repo) Create(phone string) (*User, error) {
	user := User{
		PhoneNumber: phone,
		Status:      ACTIVE,
		LanguagePreference: EN,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repo) GetByID(id uint) (*User, error) {
	var user User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repo) Update(user *User) (*User, error) {
	if err := database.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repo) Delete(id uint) error {
	return database.DB.Delete(&User{}, id).Error
}