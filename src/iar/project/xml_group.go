package project

type GroupItem struct {
	Name  string      `xml:"name"`
	File  []FileName  `xml:"file"`
	Group []GroupItem `xml:"group"` //recursively defined
}

type FileName struct {
	Name string `xml:"name"`
}
