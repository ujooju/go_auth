package handlers

import "net/http"

func FileServer() http.Handler {
	fileServer := http.StripPrefix("/static", http.FileServer(http.Dir("./static/")))
	return fileServer
}

func Init_mux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", start_page)
	mux.HandleFunc("/sign_in", sign_in)
	mux.HandleFunc("/sign_up", sign_up)
	mux.HandleFunc("/home", home_page)
	mux.Handle("/static/", FileServer())
	handler := authMiddleware(mux)
	return handler
}
