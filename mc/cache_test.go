package mc

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tt := time.Now().Add(5 * time.Second)
	t.Log(tt)
	for i := 0; i < 10; i++ {
		//tn := time.Now()
		time.Sleep(1 * time.Second)
		td := time.Since(tt)
		t.Log("过了", i+1, "秒", td)
		if time.Now().After(tt) {
			t.Log("已经过期")
		}
	}
}
