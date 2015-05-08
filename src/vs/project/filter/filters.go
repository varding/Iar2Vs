package filter

import (
	//"encoding/xml"
	"comm"
	//"fmt"
)

func Marshal(gs []*comm.Group, cfgs []*comm.Config) []byte {
	prj := new_prj()

	for i := 0; i < len(gs); i++ {
		g := gs[i]
		//fmt.Println(g.Path)
		prj.add_filter(g.Path)
		prj.add_group(g.Path, g.Files)
	}
	//fmt.Println(string(prj.toXml()))
	return prj.toXml()
}
