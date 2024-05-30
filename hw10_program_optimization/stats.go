package hw10programoptimization

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/mail"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Email string
}
type DomainStat map[string]int

var (
	ErrUnmarshalData      = errors.New("unmarshal data error")
	ErrInvalidEmail       = errors.New("invalid email error")
	ErrInappropriateEmail = errors.New("inappropriate email error")
	json                  = jsoniter.ConfigCompatibleWithStandardLibrary
)

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	scanner := bufio.NewScanner(r)
	result := make(DomainStat)
	user := &User{}
	domainSuffix := "." + domain
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), user); err != nil {
			return nil, fmt.Errorf("%s: %w", string(scanner.Bytes()), ErrUnmarshalData)
		}

		if _, err := checkEmail(user.Email, domainSuffix); err != nil {
			continue
		}

		emailDomain := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
		result[emailDomain]++
	}

	return result, nil
}

func checkEmail(email string, domainSuffix string) (bool, error) {
	if !strings.HasSuffix(email, domainSuffix) {
		return false, ErrInappropriateEmail
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return false, ErrInvalidEmail
	}

	return true, nil
}
