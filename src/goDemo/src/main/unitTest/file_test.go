package unitTest

import (
	"goDemo/src/main/inoutput"
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
