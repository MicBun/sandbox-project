package auth

import (
	"crypto/md5"
	"fmt"
	"sandbox/core"

	"gorm.io/gorm"
)

type AuthInterface interface {
	RegisterUser(email, password string) (core.User, error)
	AuthenticateUser(email, password string) (core.User, error)
}

type Auth struct {
	Db *gorm.DB
}

func CreateAuth(db *gorm.DB) AuthInterface {
	return &Auth{
		Db: db,
	}
}

func hash(plain string) string {
	data := []byte(plain)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (a *Auth) RegisterUser(email, password string) (core.User, error) {
	user := core.User{
		Email:    email,
		Password: hash(password),
	}
	err := a.Db.Save(&user).Error
	return user, err
}

func (a *Auth) AuthenticateUser(email, password string) (core.User, error) {
	var user core.User
	err := a.Db.Model(core.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, fmt.Errorf("Unable to retrieve user %w", err)
	}

	h := hash(password)
	if h != user.Password {
		return user, fmt.Errorf("Password mismatch")
	}

	return user, nil
}
