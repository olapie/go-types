package types

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

var (
	nickRegexp     = regexp.MustCompile("^[^ \n\r\t\f][^\n\r\t\f]{0,28}[^ \n\r\t\f]$")
	usernameRegexp = regexp.MustCompile("^[a-zA-Z][\\w\\.]{1,19}$")
)

type Account interface {
	AccountType() string
	String() string
}

type Username string

func (u Username) AccountType() string {
	return "username"
}

func (u Username) IsValid() bool {
	return usernameRegexp.MatchString(string(u))
}

func (u Username) Normalize() Username {
	s := string(u)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return Username(s)
}

func (u Username) String() string {
	return (string)(u)
}

type EmailAddress string

func (e EmailAddress) AccountType() string {
	return "email_address"
}

func (e EmailAddress) Normalize() EmailAddress {
	s := string(e)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return EmailAddress(s)
}

func (e EmailAddress) IsValid() bool {
	_, err := mail.ParseAddress(string(e))
	return err == nil
}

func (e EmailAddress) String() string {
	return (string)(e)
}

type Nickname string

func (n Nickname) IsValid() bool {
	return nickRegexp.MatchString(string(n))
}

func (n Nickname) Normalize() Nickname {
	s := string(n)
	s = strings.TrimSpace(s)
	return Nickname(s)
}

func ParseAccount(s string) (Account, error) {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("invalid account")
	}

	if pn, err := NewPhoneNumber(s); err == nil {
		return pn, nil
	}

	if _, err := mail.ParseAddress(s); err == nil {
		return EmailAddress(s), nil
	}

	u := Username(s)
	if u.IsValid() {
		return u, nil
	}
	return nil, errors.New("invalid account")
}

func RandomEmailAddress() EmailAddress {
	s := RandomID().Short() + "@" + RandomID().Short() + ".com"
	return EmailAddress(s)
}

func RandomPhoneNumber() *PhoneNumber {
	return MustPhoneNumber(fmt.Sprintf("+861381234%04d", RandomID()%1e4))
}

func RandomNickname() Nickname {
	return Nickname(RandomID().Short())
}
