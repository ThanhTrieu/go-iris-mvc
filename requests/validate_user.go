package requests

import (
	"regexp"
	"unicode"
)

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var rxPhone = regexp.MustCompile("^0[0-9]{9}$")
var rxUsername = regexp.MustCompile("^[a-z0-9]{3,60}$")

type Message struct {
  Username string
  Password string
	Email string
	Phone string
  Errors  map[string]string
}

func (msg *Message) Validate() bool {
  msg.Errors = make(map[string]string)

  matchEmail := rxEmail.Match([]byte(msg.Email))
  if matchEmail == false {
    msg.Errors["Email"] = "Please enter a valid email address"
  }

	matchPassword := isValidStrengPassword(msg.Password)
	if matchPassword == false {
		msg.Errors["Password"] = "At least one uppercase letter, at least one lowercase letter, at least one number, at least one special character"
	}

	matchPhone := rxPhone.Match([]byte(msg.Phone))
	if matchPhone == false {
		msg.Errors["Phone"] = "Please enter a valid phone number"
	}

	matchUsername := rxUsername.Match([]byte(msg.Username))
	if matchUsername == false {
		msg.Errors["Username"] = "Include only lowercase letters and numbers, from 3 to 60 characters"
	}

  return len(msg.Errors) == 0
}

func isValidStrengPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}