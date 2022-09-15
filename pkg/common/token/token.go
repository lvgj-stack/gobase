package token

import (
	"github.com/dgrijalva/jwt-go"
	"sync"
	"time"
)

type Config struct {
	key         string
	identityKey string
}

var (
	config = Config{"Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5", "identityKey"}
	once sync.Once
)

func Sign(identityKey string) (tokenString string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,
		"nbf":              time.Now().Unix(),
		"iat":              time.Now().Unix(),
		"exp":              time.Now().Add(1000000 * time.Hour).Unix(),
	})

	signedString, err := token.SignedString([]byte(config.key))

	return signedString, err

}