package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	eww_name := flag.String("ws", "", "*.eww file name")
	flag.Parse()
	if *eww_name != "" {
		eww_full_name := *eww_name
		if filepath.Ext(*eww_name) != ".eww" {
			eww_full_name = *eww_name + ".eww"
		}
		convert(eww_full_name)
	} else {
		fmt.Println("search *.eww file")
		search_eww()
	}
}
func search_eww() {
	fis, _ := ioutil.ReadDir(".")
	for _, f := range fis {
		if f.IsDir() == false {
			if filepath.Ext(f.Name()) == ".eww" {
				convert(f.Name())
			}
		}
	}
}
