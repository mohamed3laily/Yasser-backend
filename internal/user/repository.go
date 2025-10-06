package user

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindByPhone(phone string) (*User, error)
	Create(phone string) (*User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) (*User, error)
	UpdateLastLogin(ctx context.Context, userID uint) error
}

type repo struct{ 
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) FindByPhone(phone string) (*User, error) {
	var user User
	result := r.db.Where("phone_number = ?", phone).First(&user)
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
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repo) FindByID(id uint) (*User, error) {
	var user User
	result := r.db.Preload("District").First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *repo) Update(user *User) (*User, error) {
	if err := r.db.Save(user).Error; err != nil {
		fmt.Printf("Error fetching updated user: %v\n", err)
		return nil, err
	}
	if err := r.db.First(user, user.ID).Error; err != nil {
		fmt.Printf("Error fetching updated user: %v\n", err)
		return nil, err
	}
	return user, nil
}

func (r *repo) UpdateLastLogin(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).Model(&User{}).
		Where("id = ?", userID).
		Update("last_login", time.Now()).
		Error
}