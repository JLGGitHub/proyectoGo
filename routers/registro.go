package routers

import (
	"ejercicios/cursoUdemy/infraestructura"
	"ejercicios/cursoUdemy/models"
	"encoding/json"
	"net/http"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email esta vacio", 400)
		return
	}

	if len(t.Pass) < 4 {
		http.Error(w, "El Pass debe ser mayor a 3 caracteres", 400)
		return
	}

	_, encontrado, _ := infraestructura.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe un usuario registrado", 400)
		return
	}

	_, _, err2 := infraestructura.Insertar(t)
	if err2 != nil {
		http.Error(w, "Ocurrio un error intentando realizar el registro en Base de Datos Mongo"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
