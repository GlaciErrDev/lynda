package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/rwcarlsen/goexif/exif"
	"strings"
)

func main() {

	filepath.Walk("../../", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		switch ext {
		case ".jpg", ".jpeg":
			fmt.Println(ext)

			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			xi(f)
		}

		return nil
	})
}

func xi(f *os.File) {
	x, _ := exif.Decode(f)
	if x != nil {
		str := x.String()
		if strings.Contains(str, "ImageDescription") {

			// ImageDescription: "CROPPED and FLIPPED Statue of Liberty"
			phrase := `ImageDescription: "`
			start := strings.Index(str, phrase) + len(phrase)
			end := start + strings.Index(str[start:], `"`)
			fmt.Println(start)
			fmt.Println(end)
			fmt.Println(str[start:end])
		}
	}
}