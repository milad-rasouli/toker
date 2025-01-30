package entity

import "encoding/json"

type URL struct {
	URL    string `json:"url"`
	Detail String `json:"detail"`
}

func (u *URL) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *URL) FromString(data []byte) err {
	return json.Unmarshal(data, &u)
}
