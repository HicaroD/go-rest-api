package cookies

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetCookie(ctx echo.Context, name string) (*http.Cookie, bool) {
	value, err := ctx.Cookie(name)
	return value, err == nil
}

func SetSecureCookie(ctx echo.Context, name, value string, expirationTimeInSeconds int64) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Unix(expirationTimeInSeconds, 0)
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.Path = "/"
	ctx.SetCookie(cookie)
}
