package helpers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
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