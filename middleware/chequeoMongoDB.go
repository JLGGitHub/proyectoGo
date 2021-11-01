package middleware

import (
	"ejercicios/cursoUdemy/infraestructura"
	"net/http"
)

func ChequeoMongoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if infraestructura.ChequeoConexion() == 0 {
			http.Error(w, "Conexion infraestructura perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
