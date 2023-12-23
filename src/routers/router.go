package routers

import (
	"fmt"
	"log"
	"net/http"
	"zapys-api/src/middleware"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
)

type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	Auth   bool
}

func Configurar() {
	r := mux.NewRouter()
	c := cors.AllowAll()
	handler := c.Handler(r)
	rotas := rotasUsuarios
	for _, rota := range rotas {
		if rota.Auth {
			r.HandleFunc(rota.URI, middleware.Autenticar(rota.Funcao)).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
		}
	}

	endereco := "http://localhost:8080"
	fmt.Println("Servidor rodando em", endereco)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

//r.HandleFunc("/createuser", useradmin.CreateUser).Methods(http.MethodPost)
//r.HandleFunc("/user", useradmin.DeleteUser).Methods(http.MethodDelete)
//s.HandleFunc("/user", useradmin.GetUser).Methods(http.MethodGet)
//fmt.Println(fmt.Sprintf("Servidor rodando, http://localhost:%d", config.Porta))
//http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r)
