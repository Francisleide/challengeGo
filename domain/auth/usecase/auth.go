package usecase

import (
	"log"
	"os"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/francisleide/ChallengeGo/gateways/db/repository"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type Claims struct {
	CPF string `json:"cpf"`
	jwt.StandardClaims
}

type AuthUc struct {
	r repository.Repository
}

func NewAuthenticationUC(repo repository.Repository) AuthUc {
	return AuthUc{
		r: repo,
	}
}

func (a AuthUc) Login(CPF, secret string) bool {
	var account entities.Account
	account.CPF = CPF
	account.Secret = secret
	acc := a.r.FindOne(account.CPF)
	if reflect.DeepEqual(acc, entities.Account{}) {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(acc.Secret), []byte(secret))
	if err != nil {
		return false
	}
	return true

}

func (a AuthUc) CreateToken(CPF string, secret string) (string, error) {
	b := a.Login(CPF, secret)
	if !b {
		log.Fatal("authentication Error")
	}
	os.Setenv("ACCESS_SECRET", "asdhjkasjheee")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": CPF,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		log.Fatal("generate token error")
		return "", err
	}

	return tokenString, nil
}

func Authentication(token *jwt.Token) interface{} {
	var accessUUID string
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		accessUUID, _ = claims["user"].(string)

	}
	return accessUUID

}
