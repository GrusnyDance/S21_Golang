package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := os.Args[len(os.Args)-1]
	f := flag.Bool("f", false, "if files needed")
	d := flag.Bool("d", false, "if directories needed")
	sl := flag.Bool("sl", false, "if symbolic links needed")
	ext := flag.String("ext", "", "choose extension")
	flag.Parse()

	var noFlags bool
	if !*f && !*d && !*sl {
		noFlags = true
	}
	ParseAll(dir, *f, *d, *sl, *ext, noFlags)
}

func ParseAll(root string, f, d, sl bool, ext string, noFlags bool) {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() && info.Mode()&os.ModePerm != 0 && (noFlags || d) {
			fmt.Printf("%s\n", path)
		}
		if info.Mode()&os.ModeSymlink != 0 && info.Mode()&os.ModePerm != 0 && (noFlags || sl) {
			fmt.Printf("%s -> ", path)
			if pointee, errLink := os.Readlink(path); errLink != nil {
				fmt.Printf("[broken]\n")
			} else {
				fmt.Printf("%s\n", pointee)
			}
		}
		if info.Mode()&os.ModePerm != 0 && info.Mode()&os.ModeSymlink == 0 && !info.IsDir() && (noFlags || f) {
			if ext != "" && strings.HasSuffix(path, ext) && f {
				fmt.Printf("%s\n", path)
			} else if ext == "" {
				fmt.Printf("%s\n", path)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal()
	}
}
