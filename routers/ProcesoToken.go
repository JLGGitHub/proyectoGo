package routers

import (
	"ejercicios/cursoUdemy/infraestructura"
	"ejercicios/cursoUdemy/models"
	"errors"
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IdUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("pruebaclave")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer ")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato token invalido")
	}
	tk = splitToken[1]

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		// http.Error(w, err.Error(), 400)
		fmt.Println(claims)
		_, encontrado, _ := infraestructura.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {

			Email = claims.Email
			IdUsuario = claims.Id.Hex()
		}
		return claims, encontrado, IdUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("token invalido 2")
	}
	return claims, false, "", nil

}
