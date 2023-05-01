package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"
)

type Generator struct {
	margin     int
	numOfPixel int

	themeColor color.RGBA
	bgImg      *image.RGBA
	onePixel   *image.RGBA
}

func (gene *Generator) Init(margin int, numOfPixel int) {
	gene.margin = margin
	gene.numOfPixel = numOfPixel

	//set theme color
	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	gene.themeColor = color.RGBA{r, g, b, 255}

	gene.onePixel = image.NewRGBA(image.Rect(0, 0, 70, 70))
	draw.Draw(gene.onePixel, gene.onePixel.Bounds(), &image.Uniform{gene.themeColor}, image.ZP, draw.Src)

	//set background
	gene.bgImg = image.NewRGBA(image.Rect(0, 0, 420, 420))
	draw.Draw(gene.bgImg, gene.bgImg.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
}

func (gene *Generator) SetPixel(x int, y int) {
	if x < 0 || x >= gene.numOfPixel || y < 0 || y >= gene.numOfPixel {
		return
	}
	s_x := x*70 + gene.margin
	s_y := y*70 + gene.margin
	e_x := (x+1)*70 + gene.margin
	e_y := (y+1)*70 + gene.margin
	draw.Draw(gene.bgImg, image.Rect(s_x, s_y, e_x, e_y), gene.onePixel, image.ZP, draw.Src)
}

func (gene *Generator) Generate() {
	for x := 0; x < 3; x++ {
		for y := 0; y < 5; y++ {
			if rand.Intn(2) == 0 {
				gene.SetPixel(x, y)
				gene.SetPixel(4-x, y)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	Generator := Generator{}

	for n := 0; n < 5; n++ {
		Generator.Init(35, 5)
		Generator.Generate()

		//save
		fileName := fmt.Sprintf("Result/result_%d.png", n)
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = png.Encode(file, Generator.bgImg)
		if err != nil {
			panic(err)
		}
	}
}
