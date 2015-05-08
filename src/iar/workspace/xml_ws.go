package workspace

import (
	"encoding/xml"
)

type WorkSpace struct {
	Name     xml.Name  `xml:"workspace"`
	Projects []Project `xml:"project"`
}

type Project struct {
	Path string `xml:"path"`
}
