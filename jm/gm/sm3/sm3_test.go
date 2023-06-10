package sm3

import "testing"

func TestSM3(t *testing.T) {
	data := "this is a data"
	jm64 := To64(data)
	jm32 := To32(data)
	jm16 := To16(data)
	t.Log(jm64, len(jm64))
	t.Log(jm32, len(jm32))
	t.Log(jm16, len(jm16))
}
