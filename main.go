package main

import (
	"flag"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	var fname string
	flag.StringVar(&fname, "img", "", "Path to Image File")
	flag.Parse()

	imagef, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	img, err := jpeg.Decode(imagef)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("./output.jpg")
	if err != nil {
		log.Fatal(err)
	}

	jpeg.Encode(out, img, nil)

}
