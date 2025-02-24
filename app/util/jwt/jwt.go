package jwt

import (
	"app/app/model"
	"app/internal/logger"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// VerifyToken verifies the JWT token and returns the claims as a map or an error
func VerifyToken(raw string) (map[string]any, error) {

	token, err := jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method conforms to expected HMAC method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}

		// Use the secret key from the configuration
		secret := []byte(viper.GetString("TOKEN_SECRET_KEY"))
		return secret, nil
	})

	if err != nil {
		// Return a detailed error if token parsing fails
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	// Return a generic invalid token error if the token is not valid
	return nil, errors.New("invalid token")
}

func CreateToken(user *model.User) (string, error) {
	// Create a new token object
	tokenDurationStr := viper.GetString("TOKEN_DURATION")
	secret := []byte(viper.GetString("TOKEN_SECRET_KEY"))
	tokenDuration, err := time.ParseDuration(tokenDurationStr)
	if err != nil {
		logger.Err("[error]: %w", err)
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"sub": jwt.MapClaims{
			"id":           user.ID,
			"username":     user.Username,
			"password":     user.Password,
			"email":        user.Email,
			"first_name":   user.FirstName,
			"last_name":    user.LastName,
			"display_name": user.DisplayName,
			"role_id":      user.RoleID,
			"status":       user.Status,
		},
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(tokenDuration).Unix(),
	})
	// Generate encoded token and send it as response
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
