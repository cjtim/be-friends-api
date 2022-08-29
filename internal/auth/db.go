package auth

import (
	"errors"
	"fmt"

	r "github.com/cjtim/be-friends-api/repository"
	"golang.org/x/crypto/bcrypt"
)

func (p *lineProfile) createLineUser() (r.User, error) {
	return r.UserRepo.UpsertLine(r.User{
		Name:       p.Name,
		PictureURL: &p.Picture,
		LineUid:    &p.LineUid,
	})
}

func CreateOrgEmailPassword(u r.User) (r.User, error) {
	if u.Password == nil || *u.Password == "" {
		return r.User{}, errors.New("password cannot blank")
	}
	bs, err := bcrypt.GenerateFromPassword([]byte(*u.Password), bcrypt.MinCost)
	if err != nil {
		return r.User{}, err
	}
	hashPassword := string(bs)
	return r.UserRepo.Register(r.User{
		Name:        u.Name,
		Email:       u.Email,
		Password:    &hashPassword,
		Description: u.Description,
		PictureURL:  u.PictureURL,
		Phone:       u.Phone,
		IsOrg:       true,
	})
}

func Login(email, password string) (r.User, error) {
	if password == "" {
		return r.User{}, errors.New("password cannot blank")
	}
	u, err := r.UserRepo.GetOrgByEmailWithPassword(email)
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
