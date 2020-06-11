package main

import (
	"errors"
	"log"
	"net/http"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func errorAddrInUse() syscall.Errno {
	return syscall.EADDRINUSE
}

func errorAddrNotAvailable() syscall.Errno {
	return syscall.EADDRNOTAVAIL
}

func raiseError() error {
	time.Sleep(10 * time.Second)
	return errors.New("EADDRNOTAVAIL")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🐞"))
	})

	router.HandleFunc("/idle", func(w http.ResponseWriter, r *http.Request) {
		raiseError()
	})

	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}
