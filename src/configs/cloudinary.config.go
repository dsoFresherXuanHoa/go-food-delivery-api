package configs

import (
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

var cloudinaryInstance *cloudinary.Cloudinary

func GetCloudinaryInstance() (*cloudinary.Cloudinary, error) {
	if cloudinaryInstance == nil {
		if err := godotenv.Load(); err != nil {
			fmt.Println("Can't load .env files! Check your .env file and try again later: " + err.Error())
			return nil, err
		} else {
			var url = os.Getenv("CLOUDINARY_URL")
			if cld, err := cloudinary.NewFromURL(url); err != nil {
				fmt.Println("Can't connect to cloudinary using cloudinary API: " + err.Error())
				return nil, err
			} else {
				cloudinaryInstance = cld
				fmt.Println("Connection to cloudinary has been created!!!")
			}
		}
	}
	return cloudinaryInstance, nil
}
