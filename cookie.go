package weaver

import "net/http"

func DeleteCookie(cookie *http.Cookie, w http.ResponseWriter) {
	cookie.HttpOnly = true
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}
