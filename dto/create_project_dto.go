package dto

import "mime/multipart"

type CreateForm struct {
	Name        string                `form:"name"`
	Stack       string                `form:"stack"`
	Description string                `form:"description"`
	Live        string                `form:"live"`
	Github      string                `form:"github"`
	Image       *multipart.FileHeader `form:"image"`
	Position    int                   `form:"position"`
}
