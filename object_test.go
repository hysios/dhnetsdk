package dhnetsdk

import (
	"testing"

	"github.com/yudai/pp"
)

func TestDhObject_BoundingBox(t *testing.T) {
	obj := DhObject{}

	obj.init()
	pp.Print(obj)
}
