package vcxproj

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
)

type Project struct {
	DefaultTargets string `xml:",attr"`
	ToolsVersion   string `xml:",attr"`
	Xmlns          string `xml:"xmlns,attr"`
	ItemGroup      [4]ItemGroup
}

func new_prj() *Project {
	p := &Project{DefaultTargets: "Build", ToolsVersion: "4.0", Xmlns: "http://schemas.microsoft.com/developer/msbuild/2003"}
	i1 := &p.ItemGroup[0]
	s := "ProjectConfigurations"
	i1.Label = &s
	c1 := &ConfigurationItem{"Debug|Win32", "Debug", "Win32"}
	c2 := &ConfigurationItem{"Release|Win32", "Release", "Win32"}
	i1.ProjectConfiguration = append(i1.ProjectConfiguration, c1)
	i1.ProjectConfiguration = append(i1.ProjectConfiguration, c2)
	return p
}

func (this *Project) toXml() []byte {
	data, err := xml.MarshalIndent(this, "", "\t")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

func (this *Project) add_group(files []string) {
	for _, f := range files {
		ext := filepath.Ext(f)
		if ext == ".h" {
			this.push_clinclude(f)
		} else if ext == ".c" || ext == ".cpp" {
			this.push_clcompile(f)
		} else {
			this.push_none(f)
		}
	}
}

//internal use
func (this *Project) push_none(file string) {
	this.ItemGroup[3].push_none(file)
}

//*.h
func (this *Project) push_clinclude(file string) {
	this.ItemGroup[2].push_clinclude(file)
}

//*.c,*.cpp
func (this *Project) push_clcompile(file string) {
	this.ItemGroup[1].push_clcompile(file)
}
