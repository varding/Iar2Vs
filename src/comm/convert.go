package comm

//iar output and vs input structs

type Group struct {
	Path  string
	Files []string
}

type Config struct {
	Name    string
	Include []string
	Macros  []string
}
