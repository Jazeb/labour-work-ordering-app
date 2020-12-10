package shared;

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CreatePasswords(password string) string{ // need to specify return type here (string) in this case
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println(err)
	}
	return string(hash) // convert hash to string and return
}
