package filter

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
)

type ItemGroup struct {
	Filter    []*FilterItem
	ClCompile []*ClItem
	ClInclude []*ClItem
	None      []*ClItem
}

func (this *ItemGroup) push_clinclude(path, file string) {
	this.ClInclude = append(this.ClInclude, &ClItem{file, path})
}

func (this *ItemGroup) push_clcompile(path, file string) {
	this.ClCompile = append(this.ClCompile, &ClItem{file, path})
}

func (this *ItemGroup) push_none(path, file string) {
	this.None = append(this.None, &ClItem{file, path})
}

func (this *ItemGroup) push_filter(path string) {
	f := &FilterItem{}
	uid, _ := uuid.NewV4()
	f.UniqueIdentifier = fmt.Sprintf("{%s}", uid.String())
	f.Include = path
	this.Filter = append(this.Filter, f)
}
