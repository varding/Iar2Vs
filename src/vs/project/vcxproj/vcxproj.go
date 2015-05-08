package vcxproj

import (
	//"encoding/xml"
	"comm"
	//"fmt"
	"bytes"
)

func Marshal(prj_name, prj_guid string, gs []*comm.Group, cfgs []*comm.Config) []byte {
	prj := new_prj(cfgs)

	for i := 0; i < len(gs); i++ {
		g := gs[i]
		prj.add_item_group(g.Files)
	}

	prj.add_prj_guid(prj_name, prj_guid)
	prj.add_include(cfgs)
	prj.add_macro(cfgs)
	//fmt.Println(string(prj.toXml()))
	x := prj.toXml()
	//marshal的时候单引号被escape了，所以只能手动转回来
	return bytes.Replace(x, []byte("&#39;"), []byte("'"), -1)
}
