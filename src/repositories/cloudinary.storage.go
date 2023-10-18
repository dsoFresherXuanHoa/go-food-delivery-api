package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/utils"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type cloudinaryStorage struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryStore(cld *cloudinary.Cloudinary) *cloudinaryStorage {
	return &cloudinaryStorage{cld: cld}
}

func (s *cloudinaryStorage) UploadProductThumb(ctx context.Context, filePath string) (*uploader.UploadResult, error) {
	storageDir := os.Getenv("CLOUDINARY_STORAGE_DIR")
	if base64String, err := utils.File2Base64(filePath); err != nil {
		fmt.Println("Error while encode file to base64: " + err.Error())
		return nil, err
	} else if resp, err := s.cld.Upload.Upload(ctx, *base64String, uploader.UploadParams{
		Folder: storageDir,
	}); err != nil {
		fmt.Println("Error while upload image to cloudinary in repository: " + err.Error())
		return nil, err
	} else {
		return resp, nil
	}
}
