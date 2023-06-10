package util

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	list := []string{"a", "b", "c"}
	fmt.Println(InArray("a", list)) // true
	fmt.Println(InArray(1, list))   // false
	fmt.Println(InArray("d", list)) // false
}
