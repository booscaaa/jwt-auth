package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

//Auth .
type Auth struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
	Type    string `json:"type"`
}

//User .
type Access struct {
	ID       int64  `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//TokenAuth .
type TokenAuth struct {
	Access Access `json:"access,omitempty"`
	Exp    int64  `json:"exp,omitempty"`
	jwt.StandardClaims
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

//CreateToken .
func CreateToken(tokenAuth TokenAuth, hash string) Auth {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &tokenAuth)

	tokenstring, err := token.SignedString([]byte("hashseguro"))
	if err != nil {
		log.Fatalln(err)
	}

	return Auth{
		Token:   tokenstring,
		Refresh: hash,
		Type:    "refreshToken",
	}
}

//VerifyToken .
func VerifyToken(w http.ResponseWriter, r *http.Request) (bool, Access) {
	access := Access{}

	token, err := jwt.Parse(ExtractToken(r), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("hashseguro"), nil //crie em uma variavel de ambiente
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Not Authorized!"))
		return false, access
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapstructure.Decode(claims["user"], &access)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Not Authorized!"))
		return false, access
	}

	return true, access
}
