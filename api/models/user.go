package models

import (
	"time"

	"golang.org/x/crypto/bcrypt" // "gorm.io/gorm"
)

// "gorm.io/gorm"

type User struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	FirstName string    `gorm:"size:255;not null" json:"first_name"`
	LastName  string    `gorm:"size:255;not null" json:"last_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePasswords(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func CreateUser(user *User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func AuthenticateUser(email, password string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if !user.ComparePasswords(password) {
		return nil, nil
	}
	return &user, nil
}

func GetUserById(id string) (*User, error) {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserById(id string, user *User) (*User, error) {
	if err := db.Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUserById(id string) error {
	if err := db.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}
