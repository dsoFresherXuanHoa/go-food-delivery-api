package exceptions

import "errors"

var (
	ErrSaveProductThumbnailsToCloudinaryStorage = errors.New("cannot upload product image to cloudinary: check your request or try again later")
)
