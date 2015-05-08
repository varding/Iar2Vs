package project

import (
	"bytes"
	"comm"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	//"path/filepath"
)

func Unmarshal(file_name string) (groups []*comm.Group, cfgs []*comm.Config) {
	//abs_path := strings.Replace(path, "$WS_DIR$", ".", -1)
	//abs_path := path
	//fmt.Printf("iar prj path:%s\n", path)
	data, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	//replace iso-8859-1 by utf-8
	data = bytes.Replace(data, []byte("iso-8859-1"), []byte("utf-8"), -1)

	var prj Project
	err2 := xml.Unmarshal(data, &prj)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	return prj.extract()
}
