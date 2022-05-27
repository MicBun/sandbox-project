package auth

import (
	"sandbox/database"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAuthUser(t *testing.T) {
	database.RunTest(func(db *gorm.DB) {
		a := CreateAuth(db)
		u, err := a.RegisterUser("foo@bar.com", "securePassword")
		assert.NoError(t, err)

		assert.Equal(t, "foo@bar.com", u.Email)
		assert.NotEqual(t, "securePassword", u.Password)

		_, err = a.AuthenticateUser("foo@bar.com", "notPassword")
		assert.Error(t, err)

		u2, err := a.AuthenticateUser("foo@bar.com", "securePassword")
		assert.NoError(t, err)
		assert.Equal(t, "foo@bar.com", u2.Email)
		assert.Equal(t, u.ID, u2.ID)
	})
}
