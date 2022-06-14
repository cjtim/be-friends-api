package repository

import (
	"github.com/google/uuid"
)

type TagUserImpl struct{}

func (t *TagUserImpl) GetTagsByUserID(userID uuid.UUID) (tags []Tag, err error) {
	err = DB.Select(&tags, `SELECT t.* FROM "tag" t INNER JOIN "tag_user" tu ON tu.tag_id = t.id WHERE tu.user_id = $1`, userID)
	return
}
