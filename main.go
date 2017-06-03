package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

type pixel struct {
	r, g, b, a uint32
}

func main() {
	fmt.Println("fmt")
}

// get all images from given folder
func getImages(dir string) [][]pixel {
	var images [][]pixel

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		img := loadImage(path)
		pixels := getPixels(img)
		images = append(images, pixels)
		return nil
	})

	return images

}

//load image to operative system to work with it
func loadImage(fileName string) image.Image {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)

	if err != nil {
		log.Fatal(err)
	}
	return img
}

func getPixels(img image.Image) []pixel {
	bounds := img.Bounds()
	fmt.Println(bounds.Dx(), "x", bounds.Dy()) //debugging tool
	pixels := make([]pixel, bounds.Dx()+bounds.Dy())

	for i := 0; i < bounds.Dx()+bounds.Dy(); i++ {
		x := i * bounds.Dx()
		y := i * bounds.Dy()
		r, g, b, a := img.At(x, y).RGBA()
		pixels[i].r = r
		pixels[i].g = g
		pixels[i].b = b
		pixels[i].a = a

	}
	return pixels
}
