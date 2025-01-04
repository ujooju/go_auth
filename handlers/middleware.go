package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

var URLWhitelist = map[string]int{
	"/":        1,
	"/sign_in": 2,
	"/sign_up": 3,
	"/static":  4,
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.String())
		if _, ok := URLWhitelist[r.URL.String()]; ok || strings.HasPrefix(r.URL.String(), "/static") {
			next.ServeHTTP(w, r)
			return
		}
		_, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
