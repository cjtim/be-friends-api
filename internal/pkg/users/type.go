package users

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ProfilePic string `json:"profilePic"`

	LoginMethodID int `json:"loginMethodID"`
}
