package filter

type ClItem struct {
	Include string `xml:",attr"`
	Filter  string `xml:"Filter"`
}

type FilterItem struct {
	Include          string `xml:"Include,attr"`
	UniqueIdentifier string
}
