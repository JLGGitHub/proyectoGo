package routers

import (
	"ejercicios/cursoUdemy/infraestructura"
	"encoding/json"
	"net/http"
)

func Usuario(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if len(Id) < 1 {
		http.Error(w, "enviar parametro id", 400)
		return
	}
	perfil, err := infraestructura.BuscarUsuario(Id)
	if err != nil {
		http.Error(w, "Ocurrio un error buscando el perfil"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
