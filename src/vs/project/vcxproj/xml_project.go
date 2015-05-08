package vcxproj

import (
	"comm"
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"
)

type Project struct {
	DefaultTargets string `xml:",attr"`
	ToolsVersion   string `xml:",attr"`
	Xmlns          string `xml:"xmlns,attr"`
	ItemGroup      [4]ItemGroup
	//PropertyGroup       PropertyGroup
	PropertyGroup       []interface{}
	ItemDefinitionGroup []*ItemDefinitionGroup
	Import              [3]Import //这个很重要，不然工程属性出不来
}

func new_prj(cfgs []*comm.Config) *Project {
	p := &Project{DefaultTargets: "Build", ToolsVersion: "4.0", Xmlns: "http://schemas.microsoft.com/developer/msbuild/2003"}
	i1 := &p.ItemGroup[0]
	s := "ProjectConfigurations"
	i1.Label = &s

	for _, c := range cfgs {
		//c1 := &ConfigurationItem{"Debug|Win32", "Debug", "Win32"}
		ci := &ConfigurationItem{fmt.Sprintf("%s|Win32", c.Name), c.Name, "Win32"}
		i1.ProjectConfiguration = append(i1.ProjectConfiguration, ci)
	}

	p.Import[0].Project = "$(VCTargetsPath)\\Microsoft.Cpp.Default.props"
	p.Import[1].Project = "$(VCTargetsPath)\\Microsoft.Cpp.props"
	p.Import[2].Project = "$(VCTargetsPath)\\Microsoft.Cpp.targets"

	//c1 := &ConfigurationItem{"Debug|Win32", "Debug", "Win32"}
	//c2 := &ConfigurationItem{"Release|Win32", "Release", "Win32"}
	//i1.ProjectConfiguration = append(i1.ProjectConfiguration, c1)
	//i1.ProjectConfiguration = append(i1.ProjectConfiguration, c2)
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

////prj guid
func (this *Project) add_prj_guid(prj_name, guid string) {
	prj_guid := &PrjGUID{"Globals", guid, "Win32Proj", prj_name}
	this.PropertyGroup = append(this.PropertyGroup, prj_guid)
}

////include
func (this *Project) add_include(cfgs []*comm.Config) {
	//inc:=strings.Join(cfgs., sep)
	//"'$(Configuration)|$(Platform)'=='Debug|Win32'"
	for _, c := range cfgs {
		inc := &ProInclude{}
		inc.Condition = fmt.Sprintf("'$(Configuration)|$(Platform)'=='%s|Win32'", c.Name)
		fmt.Println(inc.Condition)
		inc.IncludePath = strings.Join(c.Include, ";")
		inc.LinkIncremental = true
		this.PropertyGroup = append(this.PropertyGroup, inc)
	}
}

////macro
func (this *Project) add_macro(cfgs []*comm.Config) {
	for _, c := range cfgs {
		macro := &ItemDefinitionGroup{}
		macro.Condition = fmt.Sprintf("'$(Configuration)|$(Platform)'=='%s|Win32'", c.Name)
		macro.ClCompile.WarningLevel = "Level3"
		macro.ClCompile.Optimization = "Disabled"
		macro.ClCompile.PreprocessorDefinitions = strings.Join(c.Macros, ";")
		this.ItemDefinitionGroup = append(this.ItemDefinitionGroup, macro)
	}
}

////ItemGroup
func (this *Project) add_item_group(files []string) {
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
