package auth

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/MDGSF/Blog/modules/models"
	"golang.org/x/crypto/scrypt"
)

// HashUserPassword encrypt user password
func HashUserPassword(password string, salt string) string {

	dk, err := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(base64.StdEncoding.EncodeToString(dk))

	return hex.EncodeToString(dk)
}

// IsUserExist judge whether username exist in db.
func IsUserExist(username string) bool {
	user := &models.TUser{}
	user.UserName = username

	if err := user.Query(); err != nil {
		return false
	}
	return true
}

// IsEmailExist judge whether email exist in db.
func IsEmailExist(email string) bool {
	user := &models.TUser{}
	user.Email = email

	if err := user.Query(); err != nil {
		return false
	}
	return true
}

// RegisterUser register user to db
func RegisterUser(username, password, email string) error {

	user := &models.TUser{}

	salt := models.GetUserSalt()
	pwd := HashUserPassword(password, salt)

	user.UserName = username
	user.PassWord = fmt.Sprintf("%s$%s", salt, pwd)
	user.Email = email
	user.NickName = username
	return user.Create()
}
