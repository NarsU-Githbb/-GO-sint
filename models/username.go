package models

type UserModel struct {
	X          bool     `json:"x"`
	Youtube    bool     `json:"linked_in"`
	Soundcloud bool     `json:"soundcloud"`
	Github     bool     `json:"github"`
	Gitea      bool     `json:"gitea"`
	Errors     []string `json:"errors,omitempty"` // Stores platform-specific errors

}
