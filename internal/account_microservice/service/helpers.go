package service

import (
	"github.com/ursulgwopp/simbir-health/internal/custom_errors"

	"crypto/sha1"
	"fmt"
	"os"
	"regexp"
)

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}

func validateName(last_name string, first_name string) error {
	if len(last_name) < 2 || len(last_name) > 30 {
		return custom_errors.ErrLastNameInvalid
	}

	if len(first_name) < 2 || len(first_name) > 30 {
		return custom_errors.ErrFirstNameInvalid
	}

	return nil
}

func validateUsername(username string) error {
	if len(username) < 3 || len(username) > 30 {
		return custom_errors.ErrUsernameInvalidLength
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_]+$`, username); !matched {
		return custom_errors.ErrUsernameInvalidCharacters
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < 4 {
		return custom_errors.ErrShortPassword
	}

	// // Check for at least one uppercase letter
	// if matched, _ := regexp.MatchString(`[A-Z]`, password); !matched {
	// 	return errors.New("password must contain at least one uppercase letter")
	// }

	// // Check for at least one lowercase letter
	// if matched, _ := regexp.MatchString(`[a-z]`, password); !matched {
	// 	return errors.New("password must contain at least one lowercase letter")
	// }

	// if matched, _ := regexp.MatchString(`[0-9]`, password); !matched {
	// 	return custom_errors.ErrPasswordWithoutDigits
	// }

	// // Check for at least one special character
	// if matched, _ := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, password); !matched {
	// 	return errors.New("password must contain at least one special character")
	// }

	return nil
}
