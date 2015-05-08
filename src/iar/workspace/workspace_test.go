package workspace

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestWs(t *testing.T) {
	fmt.Println(filepath.Abs("."))

	Parse("../../iar_xml/pid.eww")
	t.Error(" ")
}
