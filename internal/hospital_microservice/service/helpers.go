package service

import (
	"errors"
	"regexp"

	"github.com/ursulgwopp/simbir-health/internal/custom_errors"
)

func validateName(name string) error {
	if len(name) < 2 || len(name) > 50 {
		return custom_errors.ErrLastNameInvalid
	}

	return nil
}

func validateAddress(address string) error {
	if len(address) < 2 || len(address) > 100 {
		return custom_errors.ErrLastNameInvalid
	}

	return nil
}

func validatePhone(phone string) error {
	re := regexp.MustCompile(`^(?:\+7|8)?\s*\(?(\d{3})\)?[\s-]?(\d{3})[\s-]?(\d{2})[\s-]?(\d{2})$`)

	if !re.MatchString(phone) {
		return errors.New("invalid phone number format")
	}
	return nil
}
