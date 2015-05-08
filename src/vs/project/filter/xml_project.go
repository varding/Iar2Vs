package filter

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
)

//vcxproj.filters
type Project struct {
	ToolsVersion string       `xml:",attr"`
	Xmlns        string       `xml:"xmlns,attr"`
	ItemGroup    [4]ItemGroup //filter,clcompile,clinclude,none
}

func new_prj() *Project {
	return &Project{ToolsVersion: "4.0", Xmlns: "http://schemas.microsoft.com/developer/msbuild/2003"}
}

func (this *Project) toXml() []byte {
	data, err := xml.MarshalIndent(this, "", "\t")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

//
func (this *Project) add_filter(path string) {
	this.ItemGroup[0].push_filter(path)
}

func (this *Project) add_group(path string, files []string) {
	for _, f := range files {
		ext := filepath.Ext(f)
		if ext == ".h" {
			this.push_clinclude(path, f)
		} else if ext == ".c" || ext == ".cpp" {
			this.push_clcompile(path, f)
		} else {
			this.push_none(path, f)
		}
	}
}

//internal use
func (this *Project) push_none(path, file string) {
	this.ItemGroup[3].push_none(path, file)
}

//*.h
func (this *Project) push_clinclude(path, file string) {
	this.ItemGroup[2].push_clinclude(path, file)
}

//*.c,*.cpp
func (this *Project) push_clcompile(path, file string) {
	this.ItemGroup[1].push_clcompile(path, file)
}
