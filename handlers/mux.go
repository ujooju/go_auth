package handlers

import "net/http"

func FileServer() http.Handler {
	fileServer := http.StripPrefix("/static", http.FileServer(http.Dir("./static/")))
	return fileServer
}

func Init_mux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home_page)
	mux.HandleFunc("/auth", auth)
	mux.HandleFunc("/sign_in", sign_in)
	mux.HandleFunc("/sign_up", sign_up)
	mux.Handle("/static/", FileServer())
	return mux
}
