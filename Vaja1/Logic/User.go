package Logic

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
	"vaja1/DataStructures"

	"github.com/golang-jwt/jwt"
)

func EncryptPassword(password string) (pass_hash []byte, err error) {
	// Encrypt the password
	pass_hash, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	return pass_hash, nil
}

func CheckPassword(password string, pass_hash []byte) (err error) {
	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword(pass_hash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// GenerateToken generates a JWT token containing the `user_id` and the token's lifespan
func GenerateToken(user DataStructures.User) (signedToken string, err error) {
	// Generate a JWT token which includes the `user_id` and the timeout
	tokenExpireTime := time.Now().Add(time.Minute * 10)

	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = tokenExpireTime.Unix()

	// Create the token
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	signedToken, err = unsignedToken.SignedString([]byte("myVerySecretSecretKey"))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}
	return
}

// ValidateToken validates a given JWT token with the given user's id
func ValidateToken(token string) (valid bool, userID int, err error) {
	valid = false
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("myVerySecretSecretKey"), nil
	})
	if err != nil {
		return
	}

	// Check if the token has expired
	if time.Unix(int64(claims["exp"].(float64)), 0).Before(time.Now()) {
		err = fmt.Errorf("token has expired")
		return
	}

	userID = int(claims["id"].(float64))

	// The token is valid
	valid = true
	return
}

func (c *Controller) Login(user DataStructures.User) (err error) {
	// Check if the user exists in the database
	if _, err = c.db.GetUserByUsername(user.Username); err != nil {
		return
	}
	if err = CheckPassword(user.PassHash, []byte(user.PassHash)); err != nil {
		fmt.Println("Incorrect password")
		return
	}

	return c.db.Login(user)
}
