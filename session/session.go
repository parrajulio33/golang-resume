package session

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var Store *sessions.CookieStore

func Init(secret string) {
	Store = sessions.NewCookieStore([]byte(secret))
}

// var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

const sessionName = "admin_session"

func SetUser(c echo.Context, userID, email, displayName string) error {
	sess, _ := Store.Get(c.Request(), sessionName)
	sess.Values["user_id"] = userID
	sess.Values["email"] = email
	sess.Values["display_name"] = displayName
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		// Secure: true, // enable once you're on HTTPS
	}
	return sess.Save(c.Request(), c.Response())
}

func GetUserID(c echo.Context) string {
	sess, _ := Store.Get(c.Request(), sessionName)
	userID, _ := sess.Values["user_id"].(string)
	return userID
}

func GetDisplayName(c echo.Context) string {
	sess, _ := Store.Get(c.Request(), sessionName)
	name, _ := sess.Values["display_name"].(string)
	return name
}

func Clear(c echo.Context) error {
	sess, _ := Store.Get(c.Request(), sessionName)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return sess.Save(c.Request(), c.Response())
}
