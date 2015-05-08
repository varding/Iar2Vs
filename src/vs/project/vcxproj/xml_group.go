package vcxproj

type ItemGroup struct {
	Label                *string `xml:",attr"` //only ProjectConfiguration has the attr
	ProjectConfiguration []*ConfigurationItem
	ClCompile            []*Item
	ClInclude            []*Item
	None                 []*Item
}

func (this *ItemGroup) push_clinclude(file string) {
	this.ClInclude = append(this.ClInclude, &Item{file})
}

func (this *ItemGroup) push_clcompile(file string) {
	this.ClCompile = append(this.ClCompile, &Item{file})
}

func (this *ItemGroup) push_none(file string) {
	this.None = append(this.None, &Item{file})
}

//prj_guid,PropertyGroup
type PrjGUID struct {
	Label         string `xml:",attr"`
	ProjectGuid   string
	Keyword       string
	RootNamespace string
}

//include,PropertyGroup
type ProInclude struct {
	Condition       string `xml:",attr"`
	LinkIncremental bool
	IncludePath     string
}

//macros
type ItemDefinitionGroup struct {
	Condition string `xml:",attr"`
	ClCompile struct {
		PrecompiledHeader       string
		WarningLevel            string
		Optimization            string
		PreprocessorDefinitions string
	}
	Link struct {
		SubSystem                string
		GenerateDebugInformation bool
	}
}
