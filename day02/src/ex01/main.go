package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	l := flag.Bool("l", false, "num of lines")
	m := flag.Bool("m", false, "num of characters")
	w := flag.Bool("w", false, "num of words")
	flag.Parse()

	if err := checkFlags(*l, *m, *w); err != nil {
		fmt.Println(err)
		return
	}

	wg := new(sync.WaitGroup)
	for ind, val := range os.Args {
		if ind == 0 || val[0] == '-' {
			continue
		} else {
			wg.Add(1)
			go countFunc(wg, val, *l, *m, *w)
		}
	}
	wg.Wait()
}

func countFunc(wg *sync.WaitGroup, filename string, l, m, w bool) {
	defer wg.Done()

}

func checkFlags(l, m, w bool) error {
	if l && (m || w) {
		return fmt.Errorf("Wrong num of flags")
	}
	if m && (l || w) {
		return fmt.Errorf("Wrong num of flags")
	}
	if w && (l || m) {
		return fmt.Errorf("Wrong num of flags")
	}
	return nil
}
