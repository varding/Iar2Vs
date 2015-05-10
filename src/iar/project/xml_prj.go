package project

import (
	"comm"
	"encoding/xml"
	"fmt"
	//"os"
	"path/filepath"
	"strings"
)

type Project struct {
	//xml field
	Name          xml.Name            `xml:"project"`
	Version       string              `xml:"fileVersion"`
	Configuration []ConfigurationItem `xml:"configuration"`
	Group         []GroupItem         `xml:"group"`
	File          []FileName          `xml:"file"` //file可以不属于group直接属于project
	//private
	out_gs  []*comm.Group
	out_cfg []*comm.Config
}

//print functions below
func (this *Project) extract() ([]*comm.Group, []*comm.Config) {
	this.extract_configuration()
	this.extract_group("源文件", this.Group) //加个顶层的目录，这样直接隶属*ewp的<project>节点下的file也有filter

	//先把project的file解析出来，这段转换代码从this.extract_group搬过来
	files := make([]string, len(this.File))
	for i, f := range this.File {
		files[i] = f.Name
	}
	this.out_gs = append(this.out_gs, &comm.Group{"源文件", files}) //project节点下的文件直接放到"源文件"根目录

	//.c,.h文件全部改成相对路径
	for _, gs := range this.out_gs {
		for i := 0; i < len(gs.Files); i++ {
			//os.Chdir changed current dir to $proj_dir$,so use "." instead
			f := strings.Replace(gs.Files[i], "$PROJ_DIR$", ".", -1)
			p, err := filepath.Abs(f) //to abs path
			if err != nil {
				fmt.Println(err)
				continue
			}
			gs.Files[i] = p
		}
	}

	//include也要去掉$proj_dir$
	for _, cfgs := range this.out_cfg {
		for i := 0; i < len(cfgs.Include); i++ {
			//os.Chdir changed current dir to $proj_dir$,so use "." instead
			f := strings.Replace(cfgs.Include[i], "$PROJ_DIR$", ".", -1)
			p, err := filepath.Abs(f) //to abs path
			if err != nil {
				fmt.Println(err)
				continue
			}
			cfgs.Include[i] = p
		}
	}

	return this.out_gs, this.out_cfg
}

//print configuration directly
func (this *Project) print_configuration() {
	for _, c := range this.Configuration {
		cfg_name := c.Name
		fmt.Printf("configuration name :%s\n", cfg_name)
		for _, s := range c.Settings {
			s_name := s.Name
			fmt.Printf("  settings name:%s\n", s_name)
			for _, o := range s.Data.Option {
				o_name := o.Name
				fmt.Printf("    option name:%s\n", o_name)
				o_states := o.States
				sates_str := strings.Join(o_states, ",")
				fmt.Printf("    option states:%s\n", sates_str)
			}
		}
	}
}

func (this *Project) extract_configuration() {
	for _, c := range this.Configuration {
		out_cfg := &comm.Config{Name: c.Name}
		this.out_cfg = append(this.out_cfg, out_cfg)
		//fmt.Printf("configuration name :%s\n", c.Name)
		for _, s := range c.Settings {
			s_name := s.Name
			//fmt.Printf("  settings name:%s\n", s_name)
			if s_name == "ICCARM" {
				for _, o := range s.Data.Option {
					o_name := o.Name
					if o_name == "CCDefines" { //macro
						out_cfg.Macros = o.States
					} else if o_name == "CCIncludePath2" { //headers
						out_cfg.Include = o.States
					}
				}
			}
		}
	}
}

func print_indent(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
}

//print group recursively
func (this *Project) print_group(depth int, group []GroupItem) {
	for _, g := range group {
		g_name := g.Name
		print_indent(depth)
		fmt.Printf("group name :%s\n", g_name)
		for _, f := range g.File {
			print_indent(depth)
			fmt.Printf("  file name:%s\n", f.Name)
		}
		this.print_group(depth+1, g.Group)
	}
}

func (this *Project) extract_group(parent_path string, group []GroupItem) {
	for _, g := range group {
		var cur_path string
		if parent_path == "" {
			cur_path = g.Name
		} else {
			cur_path = parent_path + "\\" + g.Name
		}
		files := make([]string, len(g.File))
		for i, f := range g.File {
			files[i] = f.Name
		}
		this.out_gs = append(this.out_gs, &comm.Group{cur_path, files})
		this.extract_group(cur_path, g.Group)
	}
}
