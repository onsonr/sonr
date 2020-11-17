package file

import (
	"os"
	//"code.google.com/p/graphics-go/graphics"
)

func NewThumbnail(path string) {
	imagePath, _ := os.Open("jellyfish.jpg")
	defer imagePath.Close()
	//srcImage, _, _ := image.Decode(imagePath)

	// Dimension of new thumbnail 80 X 80
	//dstImage := image.NewRGBA(image.Rect(0, 0, 80, 80))
	// Thumbnail function of Graphics
	//graphics.Thumbnail(dstImage, srcImage)

	newImage, _ := os.Create("thumbnail.jpg")
	defer newImage.Close()
	//jpeg.Encode(newImage, dstImage, &jpeg.Options{jpeg.DefaultQuality})
}
