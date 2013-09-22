package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	readFile1()
	readFile2()
}

func readFile1() {
	f, err := os.Open("./fileinfo-ex.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	finfo, err := f.Stat()
	if err != nil {
		panic(err)
	}

	printFileInfo(&finfo)

}

func readFile2() {
	finfo, err := os.Stat("./fileinfo-ex.txt")
	if err != nil {
		panic(err)
	}
	printFileInfo(&finfo)
}

func printFileInfo(fi *os.FileInfo) {
	finfo := *fi
	//Print file info
	fmt.Printf("Name: %s\n", finfo.Name())
	fmt.Printf("Size: %s\n", strconv.FormatInt(finfo.Size(), 10))
	fmt.Printf("Mode: %s\n", finfo.Mode().String())
	fmt.Printf("ModTime: %s\n", finfo.ModTime().Format(time.UnixDate))
	fmt.Printf("IsDir: %v\n", finfo.IsDir())
}
