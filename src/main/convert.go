package main

import (
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
)

func convert(eww_file_name string) {

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

		xml_filter_str := filter.Marshal(gs, cfgs)

		//sln和工程都要用到
		prj_guid, _ := uuid.NewV4()
		xml_gs_str := vcxproj.Marshal(file_name, prj_guid.String(), gs, cfgs)

		//保存filter
		ioutil.WriteFile(fmt.Sprintf("%s.vcxproj.filters", filepath.Join(vs_path, file_name)), xml_filter_str, os.ModePerm)
		//保存proj
		ioutil.WriteFile(fmt.Sprintf("%s.vcxproj", filepath.Join(vs_path, file_name)), xml_gs_str, os.ModePerm)
	}
}
