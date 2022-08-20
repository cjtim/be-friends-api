package img

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/cjtim/be-friends-api/internal/gstorage"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UploadResp struct {
	DownloadURL string `json:"downloadURL"`
}

// PetFileUpload - Upload file to pet object and return downloadURL
// @Summary		 Upload file to pet object and return downloadURL
// @Description  Upload file to pet object and return downloadURL
// @Tags         pet
// @Produce      json
// @accept 		 mpfd
// @Security 	Bearer
// @Param		 file 	formData	file 			true	"File upload"
// @Success      200  	{object}  	pet.UploadResp
// @Failure      400  	{string}  	string
// @Failure      500  	{string}  	string
// @Router       /api/v1/pet/img [post]
func PetFileUpload(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	data, err := file.Open()
	if err != nil {
		return err
	}
	bdata, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}

	extensions := strings.Split(file.Filename, ".")
	filename := uuid.New().String()
	path := fmt.Sprintf("pets/%s.%s", filename, extensions[len(extensions)-1])
	client, err := gstorage.GetClient()
	if err != nil {
		return err
	}
	downloadURL, err := client.Upload(path, bdata)
	if err != nil {
		return err
	}
	return c.JSON(UploadResp{DownloadURL: downloadURL})
}
