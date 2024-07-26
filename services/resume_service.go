package services

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/AsentientBanana/admin/constants"
)

func AddResume(form *multipart.FileHeader) error {

	file, err := form.Open()

	if err != nil {
		return err
	}
	defer file.Close()

	os.Remove(constants.DEFAULT_RESUME)

	local_file, err := os.Create(constants.DEFAULT_RESUME)

	if err != nil {
		return err
	}
	defer local_file.Close()

	_, copy_err := io.Copy(local_file, file)

	if copy_err != nil {
		return copy_err
	}

	return nil
}
