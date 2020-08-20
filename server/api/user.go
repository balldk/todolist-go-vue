package api

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

var secretKey []byte = []byte(os.Getenv("SECRET"))

func LoginRoute(c echo.Context) error {
	// get post data
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	// Check valid user
	var password string
	row := (*db).QueryRow("SELECT password FROM users WHERE username=?", user.Username)

	err := row.Scan(&password)
	if err == nil {
		err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))
		if err != nil {
			return c.String(http.StatusUnauthorized, "Username or Password invalid")
		}
	} else if err == sql.ErrNoRows {
		return c.String(http.StatusUnauthorized, "Username or Password invalid")
	} else {
		return err
	}

	// create jwt
	token, err := createJwt(user.Username, time.Hour*1000)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func RegisterRoute(c echo.Context) error {
	// get post data
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	// insert new user
	password, _ := hashPassword(user.Password)
	_, err := (*db).Exec("INSERT INTO users(username, password) VALUES (?, ?)", user.Username, password)

	if err != nil {
		return err
	}

	// create jwt
	token, err := createJwt(user.Username, time.Hour*1000)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func UserAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		claims := &Claims{}
		token, _ := jwt.ParseWithClaims(tokenString, claims, func(tkn *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if token == nil || !token.Valid {
			return echo.ErrUnauthorized
		}

		c.Set("username", claims.Username)
		return next(c)
	}
}
