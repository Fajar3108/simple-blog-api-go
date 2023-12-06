package configs

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY  = []byte("the secret of kalimdor")
)

const (
	LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
)
