package helpers

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"golang.org/x/crypto/bcrypt"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func RandToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func  GetSessionUserID(ctx iris.Context) int64 {
	session := sessions.Get(ctx)
	userID := session.GetInt64Default("idUserSession", 0)
	return userID
}

func GetSessionUsername(ctx iris.Context) string {
	session := sessions.Get(ctx)
	username := session.GetStringDefault("usernameSession", "")
	return username
}

func GetCurrentUserRole(ctx iris.Context) int64 {
	session := sessions.Get(ctx)
	userRole := session.GetInt64Default("roleSession", 3)
	return userRole
}

func IsLoggedIn(ctx iris.Context) bool {
	id := GetSessionUserID(ctx)
	u  := GetSessionUsername(ctx)
	if u == "" || id <= 0 {
		return false
	}
	return true
}