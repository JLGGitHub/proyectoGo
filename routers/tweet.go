package routers

import (
	"ejercicios/cursoUdemy/infraestructura"
	"ejercicios/cursoUdemy/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func GrabarTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserId:  IdUsuario,
		Mensaje: mensaje.Mensaje,
	}

	_, status, err := infraestructura.InsertarTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error db "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se pudo grabar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func LeerTweet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "Debe enviar el parametro Id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return

	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "El parametro pagina debe ser un valor numerico", 400)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := infraestructura.LeerTweet(id, pag)
	if !correcto {
		http.Error(w, "Error al leer el Tweet", 400)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
