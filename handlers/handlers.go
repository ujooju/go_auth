package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("pages/index.html")
	if err != nil {
		log.Fatal("can't open index.html")
	}
	data, _ := io.ReadAll(f)
	w.Write(data)
}

func start_page(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		return
	}
	w.Write([]byte(session.String()))
}

func auth_page(w http.ResponseWriter, r *http.Request) {
	expires := time.Now().Add(time.Hour)
	fmt.Println(expires)
	newSession := http.Cookie{
		Name:    "session_id",
		Value:   "123",
		Expires: expires,
	}
	http.SetCookie(w, &newSession)
	http.Redirect(w, r, "/start_page", http.StatusFound)
}
