package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"
	"unicode"
)

type Person struct {
	name         string
	birthdayYear int
	age          int
	email        string
	phone        string
}

type PersonInterface interface {
	setName(s string) error
	setBirthdayYear(i int) error
	setEmail(s string) error
	setPhone(s string) error
}

func NewPerson() PersonInterface {
	return &Person{}
}

func (p *Person) setName(s string) error {
	if s == "" {
		return errors.New("name is empty")
	}
	if !unicode.IsUpper(rune(s[0])) {
		return errors.New("name is not a valid format")
	}
	p.name = s
	return nil
}

func (p *Person) setBirthdayYear(i int) error {
	currentYear := time.Now().Year()
	if i < 1900 || i > currentYear {
		return errors.New("invalid value")
	}
	p.age = currentYear - i
	p.birthdayYear = i
	return nil
}

func (p *Person) setEmail(s string) error {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(s) {
		return errors.New("email is not a valid format")
	}
	p.email = s
	return nil
}

func (p *Person) setPhone(s string) error {
	if len(s) == 0 {
		return errors.New("phone must not be empty")
	}

	if s[0] == '0' {
		if len(s) != 10 && len(s) != 11 {
			return errors.New("phone is not in a valid format")
		}
	} else if s[0] == '+' {
		if len(s) != 11 && len(s) != 15 {
			return errors.New("phone is not in a valid format")
		}
	} else {
		return errors.New("phone is not in a valid format")
	}
	return nil
}

func main() {
	person := NewPerson()
	_ = person.setName("Truong")
	_ = person.setEmail("truong@gmail.com")
	_ = person.setPhone("t528357408273")
	_ = person.setBirthdayYear(2003)

	fmt.Print(person)
}
