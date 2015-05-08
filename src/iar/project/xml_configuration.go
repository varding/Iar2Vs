package project

type ConfigurationItem struct {
	Name     string         `xml:"name"`
	Settings []SettingsItem `xml:"settings"` //struct tag name match xml tag name,member type name has no effect
}

type SettingsItem struct {
	Name    string       `xml:"name"`
	Version int          `xml:"archiveVersion"`
	Data    SettingsData `xml:"data"`
}

type SettingsData struct {
	Option []OptionItem `xml:"option"`
}

type OptionItem struct {
	Name   string   `xml:"name"`
	States []string `xml:"state"` //http://blog.studygolang.com/tag/xml/ -- <Interest>下棋</Interest>
}
