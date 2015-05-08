package vcxproj

type Item struct {
	Include string `xml:",attr"`
}

// type Item struct {
// 	Include string `xml:",attr"`
// }

// type Item struct {
// 	Include string `xml:",attr"`
// }

type ConfigurationItem struct {
	Include       string `xml:",attr"`
	Configuration string
	Platform      string
}
