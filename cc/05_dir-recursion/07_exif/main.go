package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/rwcarlsen/goexif/exif"
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
		tm, _ := x.DateTime()
		fmt.Println("Taken: ", tm)

		lat, long, _ := x.LatLong()
		fmt.Println("lat, long: ", lat, ", ", long)

		fmt.Println(x.String())
	}
}