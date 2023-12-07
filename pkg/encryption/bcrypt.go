package encryption

import "golang.org/x/crypto/bcrypt"

const saltCost = 14

func Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), saltCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func Check(inputPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
