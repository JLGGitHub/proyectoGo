package handlers

import (
	"ejercicios/cursoUdemy/middleware"
	"ejercicios/cursoUdemy/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.ChequeoMongoDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.ChequeoMongoDB(routers.Login)).Methods("POST")
	router.HandleFunc("/usuario", middleware.ChequeoMongoDB(middleware.ValidarJwt(routers.Usuario))).Methods("GET")

	port := ":8080"
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(port, handler))

}
