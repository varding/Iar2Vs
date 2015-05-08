package main

import (
	"comm"
	"fmt"
	"github.com/nu7hatch/gouuid"
	iar_prj "iar/project"
	iar_ws "iar/workspace"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"vs/project/filter"
	"vs/project/vcxproj"
	"vs/sln"
)

func convert(eww_file_name string) {
	var vs_sln sln.Sln

	fmt.Printf("use \"%s\" as workspace name\n", eww_file_name)

	//exe路径
	//exe_path, _ := filepath.Abs(".")
	//exe_abs_path := filepath.Dir(exe_path)

	//eww路径
	eww_abs_name, _ := filepath.Abs(eww_file_name)
	eww_path := filepath.Dir(eww_abs_name)

	//vc工程路径
	vs_path := filepath.Join(eww_path, "vs")
	os.Mkdir(vs_path, os.ModePerm)

	//解析eww文件，提取prj的路径
	prj_files := iar_ws.Unmarshal(eww_file_name)

	for _, prj_file_name := range prj_files {
		//prj文件名是相对eww文件的路径，这儿要转成绝对路径才能让readFile正确读取
		os.Chdir(eww_path)
		prj_abs_name, _ := filepath.Abs(prj_file_name)

		//当前路径改成prj文件的路径，这样group里的文件就能获取正确的abs路径
		prj_path := filepath.Dir(prj_file_name)
		fmt.Printf("iar prj path:%s\n", prj_path)
		os.Chdir(prj_path)

		//获取工程文件名
		file_name := filepath.Base(prj_file_name)
		//去除扩展名
		s := strings.Split(file_name, ".")
		file_name = s[0]

		//解析iar project
		gs, cfgs := iar_prj.Unmarshal(prj_abs_name)

		//全部转成相对路径
		convert_rel_path(gs, cfgs, vs_path)

		//生成filter文件内容
		xml_filter_str := filter.Marshal(gs, cfgs)

		//sln和工程都要用到
		prj_guid, _ := uuid.NewV4()
		//生成vcxproj文件内容
		xml_gs_str := vcxproj.Marshal(file_name, prj_guid.String(), gs, cfgs)

		//保存filter
		ioutil.WriteFile(fmt.Sprintf("%s.vcxproj.filters", filepath.Join(vs_path, file_name)), xml_filter_str, os.ModePerm)
		//保存vcxproj
		ioutil.WriteFile(fmt.Sprintf("%s.vcxproj", filepath.Join(vs_path, file_name)), xml_gs_str, os.ModePerm)

		vs_sln.AddProject(file_name, prj_guid.String(), cfgs)
	}

	eww_name := filepath.Base(eww_file_name)
	s := strings.Split(eww_name, ".")

	sln_name := fmt.Sprintf("%s.sln", filepath.Join(vs_path, s[0]))
	ioutil.WriteFile(sln_name, vs_sln.Data(), os.ModePerm)

	//创建快捷方式
	// sln_link_name := fmt.Sprintf("%s.sln", filepath.Join(eww_path, s[0]))

	// err := os.Symlink(sln_name, sln_link_name)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func convert_rel_path(gs []*comm.Group, cfgs []*comm.Config, vs_path string) {
	//所有的代码文件路径全部改成相对的，这个防止工程文件移动位置或者复制一份新的时候vs2010还是引用了原来的位置
	for _, g := range gs {
		for i := 0; i < len(g.Files); i++ {
			g.Files[i], _ = filepath.Rel(vs_path, g.Files[i])
			//fmt.Println(g.Files[i])
		}
	}

	prj_volume_name := filepath.VolumeName(vs_path)
	//include文件有选择的改路径：同一个盘符下的就改
	for _, c := range cfgs {
		for i := 0; i < len(c.Include); i++ {
			if filepath.VolumeName(c.Include[i]) == prj_volume_name {
				c.Include[i], _ = filepath.Rel(vs_path, c.Include[i])
			}
		}
	}

}
