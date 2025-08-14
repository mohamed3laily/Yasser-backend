package user

import (
	"context"
	"fmt"
	"time"
	"yasser-backend/database"
)

type Repository interface {
	FindByPhone(phone string) (*User, error)
	Create(phone string) (*User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) (*User, error)
	UpdateLastLogin(ctx context.Context, userID uint) error
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

func (r *repo) FindByID(id uint) (*User, error) {
	var user User
	result := database.DB.Preload("District").First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repo) Update(user *User) (*User, error) {
	if err := database.DB.Save(user).Error; err != nil {
		fmt.Printf("Error fetching updated user: %v\n", err)
		return nil, err
	}
	if err := database.DB.First(user, user.ID).Error; err != nil {
		fmt.Printf("Error fetching updated user: %v\n", err)
		return nil, err
	}
	return user, nil
}

func (r *repo) UpdateLastLogin(ctx context.Context, userID uint) error {
	return database.DB.WithContext(ctx).Model(&User{}).
		Where("id = ?", userID).
		Update("last_login", time.Now()).
		Error
}