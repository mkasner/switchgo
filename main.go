package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
		goroot string
		path   string
		new    string
	)

	flag.StringVar(&goroot, "goroot", "go", "goroot")
	flag.StringVar(&path, "path", "/usr/local", "path")
	flag.StringVar(&new, "new", "go-1.6.2", "new go")

	flag.Parse()
	version, err := ioutil.ReadFile(filepath.Join(path, new, "VERSION"))
	if err != nil {
		log.Fatal(err)
	}

	// delete old goroot
	err = os.Remove(filepath.Join(path, goroot))
	if err != nil {
		log.Fatal(err)
	}
	err = os.Symlink(filepath.Join(path, new), filepath.Join(path, goroot))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New go version: %s\n", string(version))
}