package line

import (
	r "github.com/cjtim/be-friends-api/repository"
)

func (p *LineProfile) CreateLineUser() (r.User, error) {
	return r.UserRepo.UpsertLine(r.User{
		Name:       p.Name,
		PictureURL: &p.Picture,
		LineUid:    &p.LineUid,
	})
}
