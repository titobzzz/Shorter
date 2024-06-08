package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/titobzzz/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main (){
	log.SetReportCaller(true)
	r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println( "starting GO API service ...")
	fmt.Println( "
	G0 API ")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}