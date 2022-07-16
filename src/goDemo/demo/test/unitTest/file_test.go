package unitTest

import (
	"goDemo/demo/internal/inoutput"
	"testing"
)

func TestFile(t *testing.T) {
	a := inoutput.A{
		"limin",
		20,
		88.8,
	}
	inoutput.Store(&a)
	var b *inoutput.A
	inoutput.ReStore(b)
}
