package entity

import "encoding/json"

type URL struct {
	URL    string `json:"url" binding:"required,url"`
	Detail string `json:"detail omitempty"`
}

func (u *URL) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *URL) FromString(data []byte) error {
	return json.Unmarshal(data, &u)
}
