package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User struct to represent user model
type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:100;not null" json:"password"`
	Role     Role
}

// Hash password with bcrypt
func Hash(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
}

// CheckPassword ...
func CheckPassword(hashedPw, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw))
}

// BeforeSave creates a hash before save user password
func (u *User) BeforeSave() error {
	hashedPw, err := Hash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPw)
	return nil
}

// SaveUser ...
func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	err := db.Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// FindUserByID ...
func (u *User) FindUserByID(db *gorm.DB, id int) (*User, error) {

	err := db.Model(User{}).Where("id = ?", id).Take(&u).Error

	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User not found")
	}

	return u, nil
}
