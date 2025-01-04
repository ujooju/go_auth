package handlers

import (
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

func auth(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
		return
	}
	w.Write([]byte(session.String()))
}

func sign_in(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		f, _ := os.Open("pages/sign_in.html")
		data, _ := io.ReadAll(f)
		w.Write(data)
		return
	}
	expires := time.Now().Add(time.Second * 15)
	newSession := http.Cookie{
		Name:    "session_id",
		Value:   "123",
		Expires: expires,
	}
	http.SetCookie(w, &newSession)
	http.Redirect(w, r, "/auth", http.StatusFound)
}

func sign_up(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		f, _ := os.Open("pages/sign_up.html")
		data, _ := io.ReadAll(f)
		w.Write(data)
		return
	}
}
