package usecase

import "fmt"

func (u *ProductUc) UploadCompleteProduct(value []byte) error {
	fmt.Println("UploadCompleteProduct hit")
	return nil
}
