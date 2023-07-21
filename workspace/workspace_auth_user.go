package workspace

type AuthorisedUser struct {
	Email        string `json:"email,omitempty"`
	Name         string `json:"name,omitempty"`
	ChipNumber   uint64 `json:"chipNumber,omitempty"`
	Organization string `json:"organization,omitempty"`
}
