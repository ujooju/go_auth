package handlers

import "net/http"

func Init_mux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home_page)
	mux.HandleFunc("/start_page", start_page)
	mux.HandleFunc("/auth", auth_page)
	return mux
}
