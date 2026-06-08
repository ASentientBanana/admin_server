package models

type Project struct {
	GormBase
	Name        string `json:"name" xml:"name"`
	Stack       string `json:"stack" xml:"stack"`
	Description string `json:"description" xml:"description"`
	Live        string `json:"live" xml:"live"`
	Github      string `json:"github" xml:"github"`
	Image       string `json:"image" xml:"image"`
	Position    int    `json:"position" xml:"position"`
}
