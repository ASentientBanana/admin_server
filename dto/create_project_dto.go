package dto

import "mime/multipart"

type CreateForm struct {
	Name        string `form:"name"`
	Stack       string `form:"stack"`
	Description string `form:"description"`
	Live        string `form:"live"`
	Github      string `form:"github"`
	Position    int    `form:"position"`
	// The image field si going to be handle manualy
	Image *multipart.FileHeader `binding:"omitempty" form:"image"`
}
