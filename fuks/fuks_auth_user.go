package fuks

import "fmt"

type AuthorisedUser struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	ChipNumber   uint64 `json:"chipNumber,omitempty"`
	Organization string `json:"organization,omitempty"`
}

func (user AuthorisedUser) GetLogName() (name string) {
	return fmt.Sprintf("%s (%s)", user.Name, user.Organization)
}
