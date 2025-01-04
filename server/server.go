package server

import (
	"net/http"

	"github.com/ujooju/go_auth/handlers"
)

func StartServer(port string) {
	server := http.Server{
		Addr:    ":" + port,
		Handler: handlers.Init_mux(),
	}
	server.ListenAndServe()
}
