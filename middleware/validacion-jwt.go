package middleware

import (
	"ejercicios/cursoUdemy/routers"
	"net/http"
)

func ValidarJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Token invalido"+err.Error(), 401)
			return
		}
		next.ServeHTTP(w, r)
	}
}
