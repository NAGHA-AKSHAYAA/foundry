package models

import (
	"foundry/initialisers"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func (user *User) Create() (*User, error) {
	results := initialisers.DB.Create(user).Error

	if results != nil {
		return nil, results
	}

	return user, nil
}

func (user *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) Update() (*User, error) {
	if err := initialisers.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	if err := initialisers.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (user *User) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
