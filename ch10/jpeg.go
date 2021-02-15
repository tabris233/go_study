// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	"image/png"
	_ "image/png" // register PNG decoder
	"os"
)

var purpose = flag.String("format", "png", "image file format")

func main() {
	flag.Parse()

	img, kind, err := image.Decode(os.Stdin)

	_ = kind
	// fmt.Println("purpose: ", *purpose, "kind: ", kind)

	if err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}

	out := os.Stdout

	switch *purpose {
	case "png":
		if err := png.Encode(out, img); err != nil {
			fmt.Fprintf(os.Stderr, "png: %v\n", err)
			os.Exit(1)
		}
	case "jpg", "jpeg":
		if err := jpeg.Encode(out, img, &jpeg.Options{Quality: 95}); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
	case "gif":
		if err := gif.Encode(out, img, nil); err != nil {
			fmt.Fprintf(os.Stderr, "gif: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "unkonw purpose image format\n")
		os.Exit(1)
	}
}
