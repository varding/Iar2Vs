package sln

import (
	"bytes"
	"comm"
	"fmt"
)

//project type {8BC9CEB8-8B4A-11D0-8D11-00A0C91BC942}
type Sln struct {
	projects []*Prj
}

type Prj struct {
	name string
	guid string
	cfg  []*comm.Config
}

func (this *Sln) AddProject(prj_name, prj_guid string, cfg []*comm.Config) {
	p := &Prj{prj_name, prj_guid, cfg}
	this.projects = append(this.projects, p)
}

func (this *Sln) Data() []byte {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Microsoft Visual Studio Solution File, Format Version 11.00\r\n# Visual Studio 2010\r\n")

	for _, p := range this.projects {
		s := fmt.Sprintf(`Project("{8BC9CEB8-8B4A-11D0-8D11-00A0C91BC942}") = "%s", "%s.vcxproj", "{%s}"`, p.name, p.name, p.guid)
		buf.WriteString(s)
		buf.WriteString("\r\nEndProject\r\n")
	}

	buf.WriteString("Global\r\n\tGlobalSection(SolutionConfigurationPlatforms) = preSolution\r\n\t\tDebug|Win32 = Debug|Win32\r\n\t\tRelease|Win32 = Release|Win32\r\n\tEndGlobalSection\r\n\tGlobalSection(ProjectConfigurationPlatforms) = postSolution\r\n")

	for _, p := range this.projects {
		for _, c := range p.cfg {
			s := fmt.Sprintf("\t\t{%s}.%s|Win32.ActiveCfg = %s|Win32\r\n", p.guid, c.Name, c.Name)
			buf.WriteString(s)

			s = fmt.Sprintf("\t\t{%s}.%s|Win32.Build.0 = %s|Win32\r\n", p.guid, c.Name, c.Name)
			buf.WriteString(s)
		}
	}
	buf.WriteString("\tEndGlobalSection\r\n\tGlobalSection(SolutionProperties) = preSolution\r\n\t\tHideSolutionNode = FALSE\r\n\tEndGlobalSection\r\nEndGlobal")
	return buf.Bytes()
}
