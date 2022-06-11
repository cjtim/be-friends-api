package auth

import (
	"errors"
	"fmt"

	r "github.com/cjtim/be-friends-api/repository"
	"golang.org/x/crypto/bcrypt"
)

func (p *LineProfile) createLineUser() (r.User, error) {
	return r.UserRepo.UpsertLine(r.User{
		Name:       p.Name,
		PictureURL: &p.Picture,
		LineUid:    &p.LineUid,
	})
}

func CreateUserEmailPassword(name, email, rawPassword string) (r.User, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.MinCost)
	if err != nil {
		return r.User{}, err
	}
	hashPassword := string(bs)
	return r.UserRepo.RegisterUser(r.User{
		Name:     name,
		Email:    &email,
		Password: &hashPassword,
	})
}

func Login(email, password string) (r.User, error) {
	u, err := r.UserRepo.GetUserByEmailWithPassword(email)
	if err != nil {
		errStr := fmt.Sprintf("no user in db, err: %s", err.Error())
		return r.User{}, errors.New(errStr)
	}

	hashedPassword := *u.Password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return r.User{}, errors.New("password not match")
	}
	return u, err
}
