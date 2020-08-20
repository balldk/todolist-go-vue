package api

import (
	"database/sql"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pr0f3ss0r-b/todolist-go-vue/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Task struct {
	TaskId    int    `json:"taskId"`
	Owner     string `json:"owner"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}
type Claims struct {
	Username string
	jwt.StandardClaims
}

var db **sql.DB = &database.DBCon

func hashPassword(passStr string) (string, error) {
	passByte := []byte(passStr)

	passHashed, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password")
		return "", err
	}
	return string(passHashed), nil
}

func createJwt(username string, expireDuration time.Duration) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
