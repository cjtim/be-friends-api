package repository

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type TagUserImpl struct{}

func (t *TagUserImpl) GetTagsByUserID(userID uuid.UUID) (tags []Tag, err error) {
	err = DB.Select(&tags, `SELECT t.* FROM "tag" t INNER JOIN "tag_user" tu ON tu.tag_id = t.id WHERE tu.user_id = $1`, userID)
	return
}

func (t *TagUserImpl) Upsert(userID uuid.UUID, ids []int) (err error) {
	newData := []string{}
	for _, tagId := range ids {
		newData = append(newData, fmt.Sprintf(`('%s', %d)`, userID, tagId))
	}
	stm := fmt.Sprintf(`
		BEGIN;

		DELETE FROM "tag_user"
		WHERE "user_id" = '%s';

		INSERT INTO "tag_user" (user_id, tag_id)
		VALUES %s;

		COMMIT;
	`, userID, strings.Join(newData, ", "))
	_, err = DB.Query(stm)
	return
}

func (t *TagUserImpl) DeleteById(userID uuid.UUID, tagID int) (err error) {
	stm := `
		DELETE FROM "tag_user"
		WHERE "user_id" = $1 AND "tag_id" = $2
	`
	_, err = DB.Query(stm, userID, tagID)
	return
}
