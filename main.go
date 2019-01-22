package main

import (
	"ccitt/ccitt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func getPixels(file io.Reader) ([][]byte, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	var pixels [][]byte
	for y := 0; y < h; y++ {
		var row []byte
		for x := 0; x < w; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			if r == 65535 && g == 65535 && b == 65535 {
				// append white
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}

		pixels = append(pixels, row)
	}

	return pixels, nil
}

// sliceDiff compares two slices returning the first index of the different
// elements pair. Returns -1 if the slices contain the same elements
func slicesDiff(s1, s2 []byte) int {
	minLen := 0

	if len(s1) < len(s2) {
		minLen = len(s1)
	} else {
		minLen = len(s2)
	}

	for i := 0; i < minLen; i++ {
		if s1[i] != s2[i] {
			return i
		}
	}

	return -1
}

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open("/home/darkrengarius/Downloads/scan2.png")
	if err != nil {
		log.Fatalf("Error opening file: %v\n", err)
	}
	defer file.Close()

	pixels, err := getPixels(file)
	if err != nil {
		log.Fatalf("Error decoding the image: %v\n", err)
	}

	encoder := &ccitt.Encoder{BlackIs1: true}

	encoded := encoder.Encode(pixels)

	preparedBytes, err := ioutil.ReadFile("/home/darkrengarius/Downloads/scan2.gr3")
	if err != nil {
		log.Fatalf("Error opening gr3 file: %v\n", err)
	}

	log.Println(encoded)
	log.Println(preparedBytes)
	diffInd := slicesDiff(encoded, preparedBytes)
	if diffInd != -1 {
		log.Fatalf("Slices differ in %v. Encoded: %v, prepared: %v\n", diffInd,
			encoded[diffInd], preparedBytes[diffInd])
	}

	if len(encoded) != len(preparedBytes) {
		log.Fatalf("Slices differ in length")
	}

	log.Println("Slices are totally equal")
}
