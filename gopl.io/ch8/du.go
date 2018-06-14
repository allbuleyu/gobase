package ch8

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

func WalkDir(dir string, fileSizes chan<- int64, depth int64)  {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			fmt.Printf("%*s%s \n", depth * 2, "", entry.Name())
			subdir := filepath.Join(dir, entry.Name())
			WalkDir(subdir, fileSizes, depth+1)
		}else {
			fmt.Printf("%*s%s %d \n", depth*2, "", entry.Name(), entry.Size())
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v \n", err)
		return nil
	}

	return entries
}
