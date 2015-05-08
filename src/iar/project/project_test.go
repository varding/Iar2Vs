package project

import (
	"fmt"
	"testing"
)

func TestProject(t *testing.T) {
	gs, cfgs := Parse("converter.ewp")

	fmt.Println("groups:")
	for _, g := range gs {
		fmt.Println(g)
	}
	fmt.Println("configs:")
	for _, c := range cfgs {
		fmt.Println(c)
	}
	t.Error(" ")
}

func TestChDir(t *testing.T) {
	// cur_path, _ := filepath.Abs(".")
	// fmt.Printf("cur path:%s\n", cur_path)
	// //change rel path to abspath
	// os.Chdir(iar_prj_path)
	// cur_path2, _ := filepath.Abs(".")
	// fmt.Printf("cur path:%s\n", cur_path2)
}
