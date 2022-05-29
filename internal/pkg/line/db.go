package line

import (
	"context"

	r "github.com/cjtim/be-friends-api/repository"
)

func (p *LineProfile) CreateLineUser() error {
	ctx := context.Background()

	lineUser, err := r.DB.LineUsers.FindFirst(
		r.LineUsers.LineUId.Equals(p.LineUid),
	).Exec(ctx)
	newUser := err == r.ErrNotFound
	if err != nil && !newUser {
		return err
	}

	var createdUser *r.UsersModel
	if newUser {
		createdUser, err = r.DB.Users.CreateOne(
			r.Users.Name.Set(p.Name),
			r.Users.LoginMethods.Link(
				r.LoginMethods.ID.Equals(2),
			),
		).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = r.DB.LineUsers.CreateOne(
			r.LineUsers.LineUId.Set(p.LineUid),
			r.LineUsers.Name.Set(p.Name),
			r.LineUsers.User.Link(
				r.Users.ID.Equals(createdUser.ID),
			),
			r.LineUsers.ProfilePic.Set(p.Picture),
		).Exec(ctx)
		return err
	}

	// Old user
	// Update profilePic
	_, err = r.DB.LineUsers.UpsertOne(
		r.LineUsers.LineUId.Equals(p.LineUid),
	).Create(
		r.LineUsers.LineUId.Set(p.LineUid),
		r.LineUsers.Name.Set(p.Name),
		r.LineUsers.User.Link(
			r.Users.ID.Equals(lineUser.UserID),
		),
		r.LineUsers.ProfilePic.Set(p.Picture),
	).Update(
		r.LineUsers.ProfilePic.Set(p.Picture),
	).Exec(ctx)

	return err
}
