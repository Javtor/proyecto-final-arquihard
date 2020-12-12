package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"golang.org/x/image/bmp"
)

type rgb struct {
	r uint8
	g uint8
	b uint8
}

const (
	N_MUESTRAS = 100
)

var (
	pcVersion  string
	t          string
	imgVersion string

	inputImgPath    = filepath.FromSlash("./img/%v.bmp")
	outputImgPath   = filepath.FromSlash("./img/inverted_%v.bmp")
	outputFileName  = "pc%v-go-%v-version%v-tratamiento%s.csv"
	csvFile         = "apilados.csv"
	metricsFilename = "metrics.csv"
)

func invert(t int, in, out string) error {
	dat, err := os.Open(in)
	if err != nil {
		return err
	}

	defer func() {
		errClose := dat.Close()
		if err == nil {
			err = errClose
		}
	}()

	img, err := bmp.Decode(dat)
	if err != nil {
		return err
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	rgbArr0 := makeArray(height, width, img)

	rgbArr := makeArray(height, width, img)

	return writeImg(t, height, width, rgbArr0, rgbArr)
}

func makeArray(height, width int, img image.Image) [][]rgb {
	ImRGB0 := [][]rgb{}

	for r := 0; r < height; r++ {
		row := []rgb{}
		for c := 0; c < width; c++ {
			rx, gx, bx, ax := img.At(c, r).RGBA()
			r, g, b, _ := rx>>8, gx>>8, bx>>8, ax>>8

			temp := rgb{
				r: uint8(r),
				g: uint8(g),
				b: uint8(b),
			}
			row = append(row, temp)
		}
		ImRGB0 = append(ImRGB0, row)
	}

	return ImRGB0
}

func writeImg(version, height, width int, rgbArr0, rgbArr [][]rgb) error {
	f, err := os.Create(filepath.Join("data", fmt.Sprintf(outputFileName, pcVersion, imgVersion, version, t)))
	if err != nil {
		return err
	}

	fCsv, err := os.OpenFile(csvFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	for n := 0; n < N_MUESTRAS; n++ {
		start := time.Now()
		switch version {
		case 1:
			for r := 0; r < height; r++ {
				for c := 0; c < width; c++ {
					rgbArr[r][c].r = 255 - rgbArr0[r][c].r
					rgbArr[r][c].g = 255 - rgbArr0[r][c].g
					rgbArr[r][c].b = 255 - rgbArr0[r][c].b
				}
			}
		case 2:
			for r := 0; r < height; r++ {
				for c := 0; c < width; c++ {
					rgbArr[r][c].r = 255 - rgbArr0[r][c].r
				}
			}

			for r := 0; r < height; r++ {
				for c := 0; c < width; c++ {
					rgbArr[r][c].g = 255 - rgbArr0[r][c].g
				}
			}
			for r := 0; r < height; r++ {
				for c := 0; c < width; c++ {
					rgbArr[r][c].b = 255 - rgbArr0[r][c].b
				}
			}

		case 3:
			for c := 0; c < width; c++ {
				for r := 0; r < height; r++ {
					rgbArr[r][c].r = 255 - rgbArr0[r][c].r
					rgbArr[r][c].g = 255 - rgbArr0[r][c].g
					rgbArr[r][c].b = 255 - rgbArr0[r][c].b
				}
			}

		case 4:
			for r := 0; r < height; r++ {
				for c := 0; c < width; c++ {
					rgbArr[r][c].r = 255 - rgbArr0[r][c].r
				}
			}

			for r := 0; r < height; r++ {
				for c := 0; c < width; c++ {
					rgbArr[r][c].g = 255 - rgbArr0[r][c].g
					rgbArr[r][c].b = 255 - rgbArr0[r][c].b
				}
			}

		case 5:
			for r := 0; r < height; r += 2 {
				for c := 0; c < width; c += 2 {
					rgbArr[r][c].r = 255 - rgbArr0[r][c].r
					rgbArr[r][c].g = 255 - rgbArr0[r][c].g
					rgbArr[r][c].b = 255 - rgbArr0[r][c].b

					rgbArr[r][c+1].g = 255 - rgbArr0[r][c+1].g
					rgbArr[r][c+1].r = 255 - rgbArr0[r][c+1].r
					rgbArr[r][c+1].b = 255 - rgbArr0[r][c+1].b

					rgbArr[r+1][c].r = 255 - rgbArr0[r+1][c].r
					rgbArr[r+1][c].g = 255 - rgbArr0[r+1][c].g
					rgbArr[r+1][c].b = 255 - rgbArr0[r+1][c].b

					rgbArr[r+1][c+1].r = 255 - rgbArr0[r+1][c+1].r
					rgbArr[r+1][c+1].g = 255 - rgbArr0[r+1][c+1].g
					rgbArr[r+1][c+1].b = 255 - rgbArr0[r+1][c+1].b
				}
			}

		}

		stop := time.Now()
		elapsed := stop.Sub(start).Nanoseconds()

		normalized := float64(elapsed) / float64(width*height)

		row := []string{pcVersion, imgVersion, strconv.FormatInt(int64(version), 10), "go", t, strconv.FormatFloat(normalized, 'f', 3, 64)}

		_, err = f.WriteString(strconv.FormatFloat(normalized, 'f', 3, 64) + "\n")
		if err != nil {
			return err
		}

		err = WriteMetrics(pcVersion, "go", imgVersion, strconv.FormatInt(int64(version), 10), t, filepath.Join("data", fmt.Sprintf(outputFileName, pcVersion, imgVersion, version, t)), metricsFilename)
		if err != nil {
			return err
		}

		writerFullCSVFile := csv.NewWriter(fCsv)
		writerFullCSVFile.Comma = ';'
		err = writerFullCSVFile.Write(row)
		if err != nil {
			return err
		}

		writerFullCSVFile.Flush()
	}
	err = f.Close()
	if err != nil {
		return err
	}
	err = fCsv.Close()
	if err != nil {
		return err
	}
	//Write new img

	fImg, err := os.Create(outputImgPath)
	if err != nil {
		return err
	}

	defer fImg.Close()

	upLeft := image.Point{0, 0}
	upRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, upRight})

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			color := color.RGBA{rgbArr[r][c].r, rgbArr[r][c].g, rgbArr[r][c].b, 255}
			img.Set(c, r, color)
		}
	}

	return bmp.Encode(fImg, img)
}

func main() {

	args := os.Args
	pcVersion = args[1]
	t = args[3]
	imgVersion = args[4]

	inputImgPath = fmt.Sprintf(inputImgPath, imgVersion)
	outputImgPath = fmt.Sprintf(outputImgPath, imgVersion)

	version, err := strconv.Atoi(args[2])
	if err != nil {
		panic(err)
	}

	err = invert(version, inputImgPath, outputImgPath)
	if err != nil {
		fmt.Println(err)
	}
}
