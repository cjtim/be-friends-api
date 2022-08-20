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
	INSERT INTO "pet_pic" (pet_id, picture_url)
	VALUES (:pet_id, :picture_url)
	`
	_, err := DB.NamedExec(stm, data)
	return err
}
