package lab6

import (
	"log"
	"os"
	"path/filepath"
)

// Task1 - Найти файл максимального размера.
func Task1(dir string) (string, int64) {
	var maxSize int64
	var maxFile string

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}

			if info.Size() > maxSize {
				maxSize = info.Size()
				maxFile = path
			}
		}
		return nil
	})

	if err != nil {
		log.Println(err)
		return "", 0
	}

	return maxFile, maxSize
}

func Task8(dir, extension string) []string {
	result := []string{}

	if extension[0] != '.' {
		extension = `.` + extension // :D
	}

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == extension {
			result = append(result, d.Name())

		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return result
}
