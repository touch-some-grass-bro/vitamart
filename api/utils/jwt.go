package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/touch-some-grass-bro/vitamart/models"
)

func IsJWTValid(jwtToken string) (valid bool, claims jwt.MapClaims, err error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(models.Config.API.JWTSecretKey), nil
	})

	if err != nil {
		return false, nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, claims, nil
	} else {
		return false, nil, err
	}
}

func GetUserEmailFromJWT(jwtToken string) (userEmail, errMsg string, status int) {
	if jwtToken == "" {
		return "", "No token provided", http.StatusBadRequest
	}
	valid, claims, err := IsJWTValid(jwtToken)
	if err != nil {
		return "", err.Error(), http.StatusBadRequest
	}
	if !valid {
		return "", "Invalid token", http.StatusUnauthorized
	}
	email, ok := claims["email"].(string)
	if !ok {
		return "", "Invalid token", http.StatusUnauthorized
	}

	return email, "", http.StatusOK
}

func SetJWTOnCookie(email string, tokenExpiry, now time.Time, w http.ResponseWriter) error {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().UTC().Add(time.Hour * 24 * 15).Unix(),
		"iat":   now.Unix(),
	})

	signedToken, err := jwtToken.SignedString([]byte(models.Config.API.JWTSecretKey))
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    signedToken,
		Path:     "/",
		Expires:  time.Now().UTC().Add(time.Hour * 24 * 15), // 15 days
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

func GetTokenJWT(email string, tokenExpiry, now time.Time, w http.ResponseWriter) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   tokenExpiry.Unix(),
		"iat":   now.Unix(),
	})

	signedToken, err := jwtToken.SignedString([]byte(models.Config.API.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
