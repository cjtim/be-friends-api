package line

import (
	r "github.com/cjtim/be-friends-api/repository"
)

func (p *LineProfile) CreateLineUser() (r.Users, error) {
	return r.UserRepo.UpsertLine(r.Users{
		Name:       p.Name,
		PictureURL: &p.Picture,
		LineUid:    &p.LineUid,
	})
}
