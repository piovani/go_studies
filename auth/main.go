package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"

	jwt "github.com/dgrijalva/jwt-go"
)

const secretKey string = "+3wkbpbY3TR2IoQfIEe4yFCy3vOxku0R4oUNhK4TUSE7euA0Ob96aZSP46FhxW97xplrE3tOleSdN117pGf52w=="

func main() {
	e := echo.New()

	e.GET("/", hello)
	e.GET("/generate", generateKey)
	e.GET("/auth", criarToken)
	e.GET("/valid", ValidarToken)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func generateKey(c echo.Context) error {
	chave := make([]byte, 64)

	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	key := base64.StdEncoding.EncodeToString(chave)

	return c.String(http.StatusOK, key)
}

func criarToken(c echo.Context) error {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = 1

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	tokenString, _ := token.SignedString([]byte(secretKey))

	fmt.Println(tokenString)

	return c.String(http.StatusOK, tokenString)
}

func ValidarToken(c echo.Context) error {
	token := extrairToken(c)

	tokenValid, _ := jwt.Parse(token, retornarSecretKey)

	if _, ok := tokenValid.Claims.(jwt.MapClaims); ok && tokenValid.Valid {
		return c.String(http.StatusUnauthorized, "Token OK")
	}

	return c.String(http.StatusUnauthorized, "Token INVÁLIDO")
}

func extrairToken(c echo.Context) string {
	tokenWithBearer := c.Request().Header.Get("Authorization")
	token := ""

	if len(strings.Split(tokenWithBearer, " ")) == 2 {
		token = strings.Split(tokenWithBearer, " ")[1]
	}

	return token
}

func retornarSecretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Metodo de assinatura errado: %v", token.Header["alg"])
	}

	return []byte(secretKey), nil
}
