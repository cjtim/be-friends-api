package repository

type PetPicImpl struct{}

type PetPic struct {
	PetID      int    `json:"pet_id" db:"pet_id"`
	PictureURL string `json:"picture_url" db:"picture_url"`
}

func (p *PetPicImpl) Add(PetID int, urls []string) error {
	data := []map[string]interface{}{}
	for _, url := range urls {
		data = append(data, map[string]interface{}{
			"pet_id":      PetID,
			"picture_url": url,
		})
	}
	stm := `
	INSERT INTO "pic_pet" (pet_id, picture_url)
	VALUES (:pet_id, :picture_url)
	`
	_, err := DB.NamedExec(stm, data)
	return err
}

func (p *PetPicImpl) DeleteAll(petID int) error {
	stm := `DELETE FROM "pic_pet" WHERE pet_id = $1`
	_, err := DB.Exec(stm, petID)
	return err
}
