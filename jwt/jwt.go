package jwt

import (
	"ejercicios/cursoUdemy/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerarTokenJwt(t models.Usuario) (string, error) {
	miClave := []byte("pruebaclave")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"nombre":    t.Nombre,
		"apellidos": t.Apellidos,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)

	if err != nil {

		return tokenStr, err
	}
	return tokenStr, nil
}
