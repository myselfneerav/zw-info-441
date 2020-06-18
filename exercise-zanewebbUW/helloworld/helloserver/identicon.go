package main

import (
	"crypto/sha256"
	"image"
	"image/color"
	"image/draw"
)

//YOU DO NOT NEED TO WRITE ANY ADDITIONAL CODE HERE
//identicon returns a unique avatar image for the provided `data`.
//The image will consist of a 4x4 grid of colors, where the
//specific colors are determined by hashing the `data` value.
func identicon(data string) image.Image {
	//hash `data` using SHA-256
	hasher := sha256.New()
	hasher.Write([]byte(data))
	hash := hasher.Sum(nil)

	//duplicate the first 16 bytes so that we have 48 bytes,
	//which will give us 16 colors we can display in a 4x4 grid
	hash = append(hash, hash[0:16]...)

	//create a new image
	img := image.NewRGBA(image.Rect(0, 0, 64, 64))

	//create a 4x4 grid of colored squares, using 3 bytes
	//for the red, green, and blue values of each color
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			i := (row * 12) + (col * 3)
			src := image.Uniform{color.RGBA{hash[i], hash[i+1], hash[i+2], 255}}
			rc := image.Rect(col*16, row*16, col*16+16, row*16+16)
			draw.Draw(img, rc, &src, image.ZP, draw.Src)
		}
	}

	return img
}
