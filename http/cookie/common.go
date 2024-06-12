package cookie

import "net/http"

func Delete(cookie *http.Cookie, w http.ResponseWriter) {
	cookie.HttpOnly = true
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}
