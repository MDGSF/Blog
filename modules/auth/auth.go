package auth

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/MDGSF/Blog/modules/models"
	"github.com/MDGSF/Blog/setting"
	"github.com/MDGSF/Blog/u"
	"github.com/astaxie/beego/session"
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
	user := &models.User{}
	user.UserName = username

	if err := user.Query(); err != nil {
		return false
	}
	return true
}

// IsEmailExist judge whether email exist in db.
func IsEmailExist(email string) bool {
	user := &models.User{}
	user.Email = email

	if err := user.Query(); err != nil {
		return false
	}
	return true
}

// RegisterUser register user to db
func RegisterUser(username, password, email string) error {

	user := &models.User{}

	salt := models.GetUserSalt()
	pwd := HashUserPassword(password, salt)

	user.UserName = username
	user.PassWord = fmt.Sprintf("%s$%s", salt, pwd)
	user.Email = email
	user.NickName = username
	return user.Create()
}

// GetUserIDFromSession get user id from session
func GetUserIDFromSession(sess session.Store) uint64 {
	userID := sess.Get("auth_user_id")
	if userID == nil {
		return 0
	}

	if id, ok := userID.(uint64); ok && id > 0 {
		return id
	}
	return 0
}

// GetUserFromSession get user from session
// @user[out]: store result in user.
// @return true: success,  false: failed.
func GetUserFromSession(user *models.User, sess session.Store) bool {
	if sess == nil {
		return false
	}

	id := GetUserIDFromSession(sess)
	if id > 0 {
		u := &models.User{ID: id}
		if u.Query() == nil {
			*user = *u
			return true
		}
	}
	return false
}

// LoginUserFromRememberCookie check user cookie
func LoginUserFromRememberCookie(user *models.User, ctx *context.Context) (success bool) {
	userName := ctx.GetCookie(setting.CookieUserName)
	if len(userName) == 0 {
		return false
	}

	defer func() {
		if !success {
			DeleteRememberCookie(ctx)
		}
	}()

	user.UserName = userName
	if err := user.Query(); err != nil {
		return false
	}

	secret := u.EncodeMd5(user.Rands + user.PassWord)
	value, _ := ctx.GetSecureCookie(secret, setting.CookieRememberName)
	if value != userName {
		return false
	}

	LoginUser(user, ctx, true)

	return true
}

// LoginUser login user
func LoginUser(user *models.User, ctx *context.Context, remember bool) {
	ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)
	ctx.Input.CruSession = beego.GlobalSessions.SessionRegenerateID(ctx.ResponseWriter, ctx.Request)
	ctx.Input.CruSession.Set("auth_user_id", user.ID)

	if remember {
		WriteRememberCookie(user, ctx)
	}
}

// LogoutUser logout user
func LogoutUser(ctx *context.Context) {
	DeleteRememberCookie(ctx)
	ctx.Input.CruSession.Delete("auth_user_id")
	ctx.Input.CruSession.Flush()
	beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
}

// WriteRememberCookie write user info to cookie
func WriteRememberCookie(user *models.User, ctx *context.Context) {
	secret := u.EncodeMd5(user.Rands + user.PassWord)
	days := 60 * 60 * 24 * setting.LoginRememberDays
	ctx.SetCookie(setting.CookieUserName, user.UserName, days)
	ctx.SetSecureCookie(secret, setting.CookieRememberName, user.UserName, days)
}

// DeleteRememberCookie delete user info in cookie
func DeleteRememberCookie(ctx *context.Context) {
	ctx.SetCookie(setting.CookieUserName, "", -1)
	ctx.SetCookie(setting.CookieRememberName, "", -1)
}
