package common_util

import (
    "golang.org/x/crypto/bcrypt"
    "net/mail"
)

func EmailValid(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

func PasswordMatch(password string, hashedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return false
    }

    return true
}
