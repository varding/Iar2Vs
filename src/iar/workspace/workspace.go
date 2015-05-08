package workspace

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

func Unmarshal(file_name string) []string {
	eww, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//replace iso-8859-1 by utf-8
	eww = bytes.Replace(eww, []byte("iso-8859-1"), []byte("utf-8"), -1)

	var ws WorkSpace
	err2 := xml.Unmarshal(eww, &ws)
	if err2 != nil {
		fmt.Println(err2)
		return nil
	}

	paths := make([]string, len(ws.Projects))
	for i, p := range ws.Projects {
		paths[i] = strings.Replace(p.Path, "$WS_DIR$", ".", -1) //
	}
	return paths
}
