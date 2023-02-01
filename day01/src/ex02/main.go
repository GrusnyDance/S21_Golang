package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Not enough args")
		return
	}

	var oldFile, newFile string
	var err error
	for ind, val := range os.Args {
		if val == "--old" && (ind+1) <= (len(os.Args)-1) {
			oldFile = os.Args[ind+1]
		}
		if val == "--new" && (ind+1) <= (len(os.Args)-1) {
			newFile = os.Args[ind+1]
		}
	}
	if oldFile == "" || newFile == "" {
		fmt.Println("I failed to get filenames")
		return
	}

	if err = checkOpenFile(oldFile, newFile); err != nil {
		fmt.Println(err)
		return
	}

	hashMap := getHashFromOldFile(oldFile)
	compare(hashMap, newFile, oldFile)
}

func compare(hashMap *map[[16]byte]bool, newFile, oldFile string) {
	newFileDescriptor, _ := os.Open(newFile)
	defer newFileDescriptor.Close()
	scanner := bufio.NewScanner(newFileDescriptor)

	for scanner.Scan() {
		hmd5 := md5.Sum(scanner.Bytes())
		if _, ok := (*hashMap)[hmd5]; !ok {
			fmt.Printf("ADDED %s\n", scanner.Text())
		} else {
			(*hashMap)[hmd5] = true
		}
	}

	oldFileDescriptor, _ := os.Open(oldFile)
	defer oldFileDescriptor.Close()
	scanner = bufio.NewScanner(oldFileDescriptor)
	for scanner.Scan() {
		hmd5 := md5.Sum(scanner.Bytes())
		if !(*hashMap)[hmd5] {
			fmt.Printf("REMOVED %s\n", scanner.Text())
		}
	}
}

func getHashFromOldFile(filename string) *map[[16]byte]bool {
	hashMap := make(map[[16]byte]bool)
	fileDescriptor, _ := os.Open(filename)
	defer fileDescriptor.Close()

	scanner := bufio.NewScanner(fileDescriptor)
	for scanner.Scan() {
		line := scanner.Bytes()
		hmd5 := md5.Sum(line)
		hashMap[hmd5] = false
	}
	return &hashMap
}

func checkOpenFile(oldFile, newFile string) error {
	var fd1, fd2 *os.File
	var err error
	if fd1, err = os.Open(oldFile); err != nil {
		return fmt.Errorf("cannot open the file")
	}
	defer fd1.Close()

	if fd2, err = os.Open(oldFile); err != nil {
		return fmt.Errorf("cannot open the file")
	}
	defer fd2.Close()
	return nil
}
