package routers

import (
	"ejercicios/cursoUdemy/infraestructura"
	"ejercicios/cursoUdemy/jwt"
	"ejercicios/cursoUdemy/models"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var model models.Usuario

	err := json.NewDecoder(r.Body).Decode(&model)

	if err != nil {
		http.Error(w, "Usuario o contraseña invalida"+err.Error(), 400)
		return
	}

	if len(model.Email) <= 0 {
		http.Error(w, "Email es requerido", 400)
		return
	}

	documento, existe := infraestructura.IntentoLogin(model.Email, model.Pass)

	if !existe {
		http.Error(w, "Usuario o contraseña invalidos", 400)
		return
	}

	jwtkey, err := jwt.GenerarTokenJwt(documento)

	if err != nil {
		http.Error(w, "No se pudo generar el token "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtkey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resp)
}
