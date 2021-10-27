package common

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math"
	"os"

	"github.com/sonr-io/core/internal/fs"
)

// NewThumbnail creates a new thumbnail for the given file
func NewThumbnail(path string, mime *MIME, ch chan *Thumbnail) {
	if fs.Exists(path) {
		if mime.IsImage() {
			logger.Info("Found Image, Generating Thumbnail")
			data, err := genImageThumbnail(path)
			if err == nil {
				ch <- &Thumbnail{
					Buffer: data,
					Mime:   mime,
				}
				return
			}
		} else {
			logger.Info("Thumbnail exists at path, Building Thumbnail")
			tbuf, err := os.ReadFile(path)
			if err != nil {
				logger.Error("%s - Failed to read thumbnail path", err)
				ch <- nil
				return
			}
			ch <- &Thumbnail{
				Buffer: tbuf,
				Mime:   mime,
			}
			return
		}
	}
	logger.Warn("Thumbnail does not exist at path, skipping...")
	ch <- nil
}

func genImageThumbnail(path string) ([]byte, error) {
	// Open Image
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	srcImage, _, _ := image.Decode(f)

	// Dimension of new thumbnail 240 X 300
	dstImage := resizeImage(srcImage, 240, 300)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, dstImage, &jpeg.Options{Quality: 75})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func resizeImage(img image.Image, length int, width int) image.Image {
	//truncate pixel size
	minX := img.Bounds().Min.X
	minY := img.Bounds().Min.Y
	maxX := img.Bounds().Max.X
	maxY := img.Bounds().Max.Y
	for (maxX-minX)%length != 0 {
		maxX--
	}
	for (maxY-minY)%width != 0 {
		maxY--
	}
	scaleX := (maxX - minX) / length
	scaleY := (maxY - minY) / width

	imgRect := image.Rect(0, 0, length, width)
	resImg := image.NewRGBA(imgRect)
	draw.Draw(resImg, resImg.Bounds(), &image.Uniform{C: color.White}, image.ZP, draw.Src)
	for y := 0; y < width; y += 1 {
		for x := 0; x < length; x += 1 {
			averageColor := getImageAverageColor(img, minX+x*scaleX, minX+(x+1)*scaleX, minY+y*scaleY, minY+(y+1)*scaleY)
			resImg.Set(x, y, averageColor)
		}
	}
	return resImg
}

func getImageAverageColor(img image.Image, minX int, maxX int, minY int, maxY int) color.Color {
	var averageRed float64
	var averageGreen float64
	var averageBlue float64
	var averageAlpha float64
	scale := 1.0 / float64((maxX-minX)*(maxY-minY))

	for i := minX; i < maxX; i++ {
		for k := minY; k < maxY; k++ {
			r, g, b, a := img.At(i, k).RGBA()
			averageRed += float64(r) * scale
			averageGreen += float64(g) * scale
			averageBlue += float64(b) * scale
			averageAlpha += float64(a) * scale
		}
	}

	averageRed = math.Sqrt(averageRed)
	averageGreen = math.Sqrt(averageGreen)
	averageBlue = math.Sqrt(averageBlue)
	averageAlpha = math.Sqrt(averageAlpha)

	averageColor := color.RGBA{
		R: uint8(averageRed),
		G: uint8(averageGreen),
		B: uint8(averageBlue),
		A: uint8(averageAlpha)}

	return averageColor
}

func imgToBytes(img image.Image) []byte {
	var opt jpeg.Options
	opt.Quality = 80

	buff := bytes.NewBuffer(nil)
	err := jpeg.Encode(buff, img, &opt)
	if err != nil {
		log.Fatal(err)
	}

	return buff.Bytes()
}
